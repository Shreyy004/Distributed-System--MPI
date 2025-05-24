#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpi.h>

void send_to_many (int num_procs);
void receive_from_one (int rank);
void await_request (int rank, MPI_Request *request);
int main (int argc, char **argv)
{
int rank, num_procs;
MPI_Init (&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
srandom(time(NULL)+rank);
printf("%d hi(count=%d)\n", rank, num_procs);
if (rank==0) { send_to_many(num_procs); }
else { receive_from_one(rank); }
printf("%d: goodbye\n", rank);
MPI_Finalize();

}

void send_to_many (int num_procs)
{
long random_value, send_value;
random_value = random()/(RAND_MAX/100);
MPI_Request send_request;
for (int i=1; i<num_procs; i++)
{
send_value = random_value + 2*i;
MPI_Isend((void *) &send_value, 1, MPI_LONG, i, 1, MPI_COMM_WORLD, &send_request);
await_request(0, &send_request);
printf("0: sending %ld to %d\n", send_value, i);
}
}

void receive_from_one (int rank)
{
long recv_value;
MPI_Request receive_request;
MPI_Irecv((void *) &recv_value, 1, MPI_LONG, 0, 1, MPI_COMM_WORLD, &receive_request);
await_request (rank, &receive_request);
printf("%d: received %ld\n", rank, recv_value);

}

void await_request(int rank, MPI_Request *request)
{
int wait_count=0;
int flag=0;
MPI_Status status;
do
{
wait_count++;
MPI_Test (request, &flag, &status);
} while (!flag);
printf("%d: tested %d times\n", rank, wait_count);
}
