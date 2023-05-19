from mrjob.job import MRJob, MRStep


class DiaMinMax(MRJob):
    def mapper(self, key, line):
        company, price, date = line.split(",")
        if line.startswith("Company"):  # skip the header line
            pass
        else:
            yield company, (date, price)

    def reducer(self, company, data):
        data_list = [x for x in data]
        lowest_date = min(data_list, key=lambda x: float(x[1]), default=0)[0]
        highest_date = min(data_list, key=lambda x: float(x[1]), default=0)[0]

        yield company, [lowest_date, highest_date]


class AccionCreciente(MRJob):
    def mapper(self, key, line):
        company, price, date = line.split(",")
        if line.startswith("Company"):  # skip the header line
            pass
        else:
            yield company, (date, price)

    def reducer(self, company, data):
        data_list = [x for x in data]
        dates_sorted = sorted(data_list, key=lambda x: x[0])
        new_stonks = [dates_sorted[i][0] for i in range(1, len(dates_sorted)) if float(dates_sorted[i][1]) >= float(dates_sorted[0][1])]
        if len(new_stonks):
            yield company, "Acciones crecientes o estables"




class BlackFriday(MRJob):
    def mapper(self, key, line):
        company, price, date = line.split(",")
        if line.startswith("Company"):  # skip the header line
            pass
        else:
            yield None, (price, date)

    def reducer(self, _, data):
        sums = {}
        for price, date in data:
            if date not in sums:
                sums[date] = 0
            sums[date] += float(price)
        key, value = min(sums.items(), key=lambda x: x[1])
        yield key, value


if __name__ == "__main__":
    print(">>>>>>>>>> Dia con valor minimo y maximo para cada empresa")
    DiaMinMax.run()

    print(">>>>>>>>>> Empresas con acciones iguales o mayores")
    AccionCreciente.run()

    print(">>>>>>>>>> Dia con precios mas bajos")
    BlackFriday.run()
