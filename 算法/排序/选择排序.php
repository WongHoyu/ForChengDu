<?php

/**
 * @intro
 * User: HuangHaoyu
 * Email: haoyuhuang@sfmail.sf-express.com
 * Date: 2021/12/14
 * Time: 7:14 下午
 */
class SelectSort {
    public function sort($nums) {
        if(empty($nums) || count($nums) == 1) {
            return $nums;
        }

        $length = count($nums);
        for($i = 0; $i < $length - 1; $i++) {
            $minIndex = $i;
            for($j = $i + 1; $j < $length; $j++) {
                if($nums[$j] < $nums[$minIndex]) {
                    $minIndex = $j;
                }
            }

            list($nums[$i], $nums[$minIndex]) = [$nums[$minIndex], $nums[$i]];
        }

        return $nums;
    }
}

$selectSort = new SelectSort();
print_r($selectSort->sort([5,1,1,2,0,0]));