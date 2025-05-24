int main (int argc, char **argv)
{
int num_process;
int rank;

MPI_Init (&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_process);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
printf("number of processes= %d\n", num_process);
printf("rank of process= %d\n", rank);
MPI_Finalize();

}
