#include "mpi.h"
#include <stdio.h>
#include <stdlib.h>
#define SIZE 5

int main (int argc, char *argv[])
{
int numtasks, rank, sendcount, recvcount, source;
float sendbuf[SIZE][SIZE] = {
  {1.0, 2.0, 3.0, 4.0},
  {5.0, 6.0, 7.0, 8.0},
  {9.0, 10.0, 11.0, 12.0},
  {13.0, 14.0, 15.0, 16.0},
  {17.0,18.0,19.0,20.0} };
float recvbuf[SIZE][SIZE];

MPI_Init(&argc,&argv);
MPI_Comm_rank(MPI_COMM_WORLD, &rank);
MPI_Comm_size(MPI_COMM_WORLD, &numtasks);

if (numtasks == SIZE) {
  source = 4;
  sendcount = SIZE;
  recvcount = SIZE;
  MPI_Gather(sendbuf[rank],sendcount,MPI_FLOAT,recvbuf,recvcount,
             MPI_FLOAT,source,MPI_COMM_WORLD);
  if (rank == source) {
            printf("Results gathered at rank %d:\n", rank);
            for (int i = 0; i < SIZE; i++) {
                printf("Row %d: %f %f %f %f\n", i, recvbuf[i][0], recvbuf[i][1], recvbuf[i][2], recvbuf[i][3]);
            }
        }
  }
else
  printf("Must specify %d processors. Terminating.\n",SIZE);
MPI_Finalize();
}

