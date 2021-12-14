<?php

/**
 * @intro
 * User: HuangHaoyu
 * Email: haoyuhuang@sfmail.sf-express.com
 * Date: 2021/12/14
 * Time: 6:45 下午
 */
class BubbleSort {
    public function sort($nums) {
        if(empty($nums) || count($nums) == 1) {
            return $nums;
        }

        $length = count($nums) - 1;
        for($i = 0; $i <= $length; $i ++) {
            for($j = 1; $j <= $length - $i; $j++) {
                if($nums[$j] < $nums[$j - 1]) {
                    list($nums[$j - 1], $nums[$j]) = [$nums[$j], $nums[$j - 1]];
                }
            }
        }

        return $nums;
    }
}

$bubbleSort = new BubbleSort();
print_r($bubbleSort->sort([5,1,1,2,0,0]));