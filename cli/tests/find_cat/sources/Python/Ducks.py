#!/usr/bin/env python3

from math import exp, sqrt


def frange(start: float, end: float = None, inc: float = 1.0):
    range_list: list[float] = []
    if end is None:
        end = start
        start = 0.0

    value = start
    while (inc > 0. and value < end) or (inc < 0. and value > end):
        range_list.append(value)
        value = start + range_list.__len__() * inc
    return range_list


class Duck:
    max = 100

    def __init__(self, a: float, test: bool = False):
        self.a = a
        self.interval = frange(0,  Duck.max, 0.001)
        self.esp = 0
        self.std_dev = 0
        self.test = test

    @staticmethod
    def percent_back(const: float, time_snd: int):
        return sum(10 * Duck.probability_density(const, i) for i in frange(0, time_snd, 0.001)) / 100

    @staticmethod
    def time_back(a: float, p: float):
        res = 0
        for t in frange(0, Duck.max, 0.01):
            res += Duck.probability_density(a, t)
            if res > p:
                return t
        raise ValueError

    @staticmethod
    def probability_density(a: float, t: float):
        return a * exp(-t) + (4 - 3 * a) * exp(-2 * t) + (2 * a - 4) * exp(-4 * t)

    @staticmethod
    def variance(esp: float, a: float, t: float):
        return pow((t - esp), 2) * (Duck.probability_density(a, t) / 10)

    @staticmethod
    def esperance(const: float, interval: list):
        return sum(time * (Duck.probability_density(const, time) / 10) for time in interval) / interval[-1]

    @staticmethod
    def standard_deviation(const: float, esp: float, interval):
        return sqrt(sum(Duck.variance(esp, const, time) for time in interval) / interval[-1])

    def print(self):
        print("Average return time: %0.0fm %0.02ds" % divmod(round(self.esp * 60), 60))
        print("Standard deviation: %.3f" % self.std_dev)
        print("Time after which 50%% of the ducks are back: %dm %02ds" % divmod(self.time_back(self.a, 50) * 60, 60))
        print("Time after which 99%% of the ducks are back: %dm %02ds" % divmod(self.time_back(self.a, 99) * 60, 60))
        print("Percentage of ducks back after 1 minute: %.1f%%" % (self.percent_back(self.a, 1)))
        print("Percentage of ducks back after 2 minutes: %.1f%%" % self.percent_back(self.a, 2))

    def run(self):
        self.esp = self.esperance(self.a, self.interval)
        self.std_dev = self.standard_deviation(self.a, self.esp, self.interval)
        return self
