<?php

/**
 * @intro
 * User: HuangHaoyu
 * Email: haoyuhuang@sfmail.sf-express.com
 * Date: 2021/12/14
 * Time: 7:07 下午
 */
class InsertSort {
    public function sort($nums) {
        if(empty($nums) || count($nums) == 1) {
            return $nums;
        }

        $length = count($nums);
        for($i = 0; $i < $length - 1; $i++) {
            for($j = $length - 1; $j > 0; $j--) {
                if($nums[$j] < $nums[$j - 1]) {
                    list($nums[$j], $nums[$j - 1]) = [$nums[$j - 1], $nums[$j]];
                }
            }
        }

        return $nums;
    }
}

$insertSort = new InsertSort();
print_r($insertSort->sort([5,1,1,2,0,0]));