<?php

class quickSort {
    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function sortArray($nums) {
        if(empty($nums) || count($nums) == 1) {
            return $nums;
        }

        $this->_quickSort($nums, 0, count($nums) - 1);
        return $nums;
    }

    private function _quickSort(&$nums, $left, $right) {
        if($left >= $right) {
            return ;
        }

        $pivot = $this->_partition($nums, $left, $right);
        $this->_quickSort($nums, $left, $pivot - 1);
        $this->_quickSort($nums, $pivot + 1, $right);
    }

    private function _partition(&$nums, $left, $right) {
        $pivot = $nums[$left];

        while($left < $right) {
            while($left < $right && $nums[$right] >= $pivot) {
                $right--;
            }
            $this->_swap($nums, $left, $right);

            while($left < $right && $nums[$left] <= $pivot) {
                $left++;
            }
            $this->_swap($nums, $left, $right);
        }

        return $left;
    }

    private function _swap(&$nums, $left, $right) {
        list($nums[$left], $nums[$right]) = [$nums[$right], $nums[$left]];
    }
}

$solution = new quickSort();