<?php

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */

class ListNode {
    public $val = 0;
    public $next = null;

    function __construct($val = 0, $next = null) {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution {

    /**
     * @param ListNode $list1
     * @param ListNode $list2
     * @return ListNode
     */
    function mergeTwoLists1($list1, $list2) {
        if($list1 === null || $list2 === null) {
            return $list1 ?: $list2;
        }

        $newList = new ListNode();
        $point = $newList;
        while($list1 !== null && $list2 !== null) {
            if($list1->val < $list2->val) {
                $point->next = $list1;
                $list1 = $list1->next;
            } else {
                $point->next = $list2;
                $list2 = $list2->next;
            }

            $point = $point->next;
        }

        $point->next = $list1 === null ? $list2 : $list1;

        return $newList->next;
    }

    /**
     * @param ListNode $list1
     * @param ListNode $list2
     * @return ListNode
     */
    function mergeTwoLists($list1, $list2) {
        if($list1 === null || $list2 === null) {
            return $list1 ?: $list2;
        }

        if($list1->val < $list2->val) {
            $list1->next = $this->mergeTwoLists($list1->next, $list2);
            return $list1;
        } else {
            $list2->next = $this->mergeTwoLists($list1, $list2->next);
            return $list2;
        }
    }
}