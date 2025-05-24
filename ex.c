/*
Problem Statement:

    Each process holds a temperature reading from a sensor.
    We want to find the maximum recorded temperature across all processes and ensure every process knows it.
*/

#include <mpi.h>
#include <stdio.h>

int main(int argc, char* argv[]) {
    int rank, size;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int local_temperature = (rank + 1) * 10;  // Different temperatures per process
    int max_temperature;

    // Find maximum temperature among all processes
    MPI_Allreduce(&local_temperature, &max_temperature, 1, MPI_INT, MPI_MAX, MPI_COMM_WORLD);

    // Each process prints the same result
    printf("Process %d: Maximum temperature recorded is %d°C\n", rank, max_temperature);

    MPI_Finalize();
    return 0;
}

