class Solution:
    def count_pairs(num1: list[int], num2: list[int]) -> int:
        def bs(arr: list[int], target: int, lo: int) -> int:
            l, r = lo, len(arr) - 1
            while l < r:
                m = (l + r) // 2
                if arr[m] == target:
                    return m
                elif arr[m] > target:
                    r = m
                else:
                    l = m + 1
            return m

        d = [num1[i] - num2[i] for i in range(len(num1))]
        d.sort()

        ans = 0
        for i, n in enumerate(d):
            if n >= 0:
                break
            highest_num = bs(d, -n, i + 1)
            ans += len(num1) - highest_num
        return ans


class Solution:
    def count_pairs(num1: list[int], num2: list[int]):
        d = [num1[i] - num2[i] for i in range(len(num1))]
        d.sort()
        return sum(len(num1) - bisect_right(d, -v, lo=i + 1) for i, v in enumerate(d))
