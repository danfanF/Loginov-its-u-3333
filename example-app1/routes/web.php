<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('index', ['products' => [['name'=>  'Orange', "cost" => "50000",  "amount" => "27",],
    ['name'=>  'Banana', "cost" => "50000",  "amount" => "12",],
    ['name'=>  'Bread', "cost" => "50000",  "amount" => "0",]],
"zhora" => "сделал @Company Zhora"]);
});
