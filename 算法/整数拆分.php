<?php

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function integerBreak($n) {
        if($n < 4) {
            return $n - 1;
        }

        $res = 1;
        while($n > 4) {
            $res *= 3;
            $n -= 3;
        }

        return $n == 0 ? $res : $res * $n;
    }
}

$solution = new Solution();
echo $solution->integerBreak(4);