<?php

/**
 * @intro
 * User: HuangHaoyu
 * Email: haoyuhuang@sfmail.sf-express.com
 * Date: 2021/12/14
 * Time: 7:35 下午
 */
class HeapSort {

    private $sortNums = [];

    public function sort($nums) {
        $this->sortNums = $nums;

        $length = count($nums) - 1;
        $beginIndex = ($length - 1) >> 1;//左下角的子节点 -- 第一个没有子节点的子节点, 如果从最后一个子节点开始，会有一堆无用的调整。
        //堆化 -- 含义就是大体地调整数组，父节点一定大于其任何子节点的值
        for($i = $length; $i >= 0; $i--) {
            $this->maxHeapify($beginIndex, $length);
        }

        //调整堆 -- 含义就是先将堆化后的数组第一个元素调到最后的子节点上，然后往复调整，最后达到大大顶堆/小顶堆的效果
        for($i = $length; $i > 0; $i--) {
            $this->_swap($this->sortNums, 0, $i);
            $this->maxHeapify(0, $i - 1);
        }

        return $this->sortNums;
    }

    /**
     * @param int $index 当前节点
     * @param int $length 节点的边界
     */
    private function maxHeapify($index, $length) {
        $left = ($index << 1) + 1;//左子节点
        $right = $left + 1;//右子节点
        $maxIndex = $left;

        if($left > $length) {
            return;
        }

        if($right <= $length && $this->sortNums[$right] > $this->sortNums[$left]) {
            $maxIndex = $right;
        }

        if($this->sortNums[$index] < $this->sortNums[$maxIndex]) {
            $this->_swap($this->sortNums, $index, $maxIndex);
            $this->maxHeapify($maxIndex, $length);
        }
    }

    private function _swap(&$nums, $left, $right) {
        list($nums[$left], $nums[$right]) = [$nums[$right], $nums[$left]];
    }
}

$heapSort = new HeapSort();
print_r($heapSort->sort([5,4,3,2,1]));