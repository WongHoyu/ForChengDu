<?php
/**
 * @intro
 * User: HuangHaoyu
 * Email: haoyuhuang@sfmail.sf-express.com
 * Date: 2021/12/12
 * Time: 4:53 下午
 */

trait Singleton {

    private static $instance = null;

    private function __construct() {

    }

    private function __sleep() {

    }

    private function __wakeup() {
    }

    private function __clone() {

    }

    public function getInstance() {
        if(isset(self::$instance)) {
            return self::$instance;
        }

        return new static();
    }
}