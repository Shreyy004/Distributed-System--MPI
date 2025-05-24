#include "mpi.h"
#include <stdio.h>
#include <stdlib.h>
#define SIZE 5

int main (int argc, char *argv[])
{
int numtasks, rank, sendcount, recvcount, source;
float sendbuf[SIZE][SIZE] = {
  {1.0, 2.0, 3.0, 4.0, 100.0},
  {5.0, 6.0, 7.0, 8.0, 200.0},
  {9.0, 10.0, 11.0, 12.0,  300.0},
  {13.0, 14.0, 15.0, 16.0, 400.0},
  {17.0,18.0,19.0,20.0,500.0} };
float recvbuf[SIZE];
float avg[SIZE];
float recvb[SIZE];

MPI_Init(&argc,&argv);
MPI_Comm_rank(MPI_COMM_WORLD, &rank);
MPI_Comm_size(MPI_COMM_WORLD, &numtasks);

if (numtasks == SIZE) {
  source = 0;
  sendcount = SIZE;
  recvcount = SIZE;
  MPI_Scatter(sendbuf,sendcount,MPI_FLOAT,recvbuf,recvcount,
             MPI_FLOAT,source,MPI_COMM_WORLD);

  printf("rank= %d  Results: %f %f %f %f %f\n",rank,recvbuf[0],
         recvbuf[1],recvbuf[2],recvbuf[3],recvbuf[4]);
         
  avg[rank]=(recvbuf[0]+recvbuf[1]+recvbuf[2]+recvbuf[3]+recvbuf[4])/SIZE;
  MPI_Gather(&avg[rank],1,MPI_FLOAT,recvb,1,
             MPI_FLOAT,source,MPI_COMM_WORLD);
  if(rank==source){
  printf("Averages: %f %f %f %f %f\n",recvb[0],recvb[1],recvb[2],recvb[3],recvb[4]);
  }
  }
else
  printf("Must specify %d processors. Terminating.\n",SIZE);
MPI_Finalize();
}
