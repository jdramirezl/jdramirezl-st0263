from mrjob.job import MRJob, MRStep
import sys


class SalarioSE(MRJob):
    def mapper(self, key, line):
        empleado, sector, salario, anio = line.split(",")
        if line.startswith("idemp"):  # skip the header line
            pass
        else:
            yield sector, salario

    def reducer(self, sector, salarios):
        total, suma = 0, 0
        for salario in salarios:
            suma += float(salario)
            total += 1
        yield sector, suma / total


class SalariosEM(MRJob):
    def mapper(self, key, line):
        empleado, sector, salario, anio = line.split(",")
        if line.startswith("idemp"):  # skip the header line
            pass
        else:
            yield empleado, salario

    def reducer(self, empleado, salarios):
        total, suma = 0, 0
        for salario in salarios:
            suma += float(salario)
            total += 1
        yield empleado, suma / total


class SectorEM(MRJob):
    def mapper(self, key, line):
        empleado, sector, salario, anio = line.split(",")
        if line.startswith("idemp"):  # skip the header line
            pass
        else:
            yield empleado, sector

    def reducer(self, empleado, sectores):
        yield empleado, len(set(sectores))


if __name__ == "__main__":
    print(">>>>>>>>>> Salarios por sector")
    SalarioSE.run()

    print(">>>>>>>>>> Salarios por empleado")
    SalariosEM.run()

    print(">>>>>>>>>> Sectores por empleado")
    SectorEM.run()
