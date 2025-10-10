from heapq import heappop, heappush


class MinMaxHeap:
    min_heap = []
    max_heap = []

    def insert(self, num):
        if len(self.max_heap) == 0:
            heappush(self.max_heap, -int(num))
        elif num < -self.max_heap[0]:
            heappush(self.max_heap, -int(num))
        elif num >= -self.max_heap[0]:
            heappush(self.min_heap, num)

    def sort(self):
        if len(self.max_heap) > len(self.min_heap) + 1:
            heappush(self.min_heap, -heappop(self.max_heap))
        elif len(self.max_heap) < len(self.min_heap):
            heappush(self.max_heap, -heappop(self.min_heap))

    def median(self):
        if (len(self.max_heap) + len(self.min_heap)) % 2 == 0:
            return (-self.max_heap[0] + self.min_heap[0]) / 2
        else:
            return float(-self.max_heap[0])


def runningMedian(a):
    heap = MinMaxHeap()
    medians = []
    for i in range(len(a)):
        heap.insert(a[i])
        heap.sort()
        medians.append(round(heap.median(), 1))
    return medians
