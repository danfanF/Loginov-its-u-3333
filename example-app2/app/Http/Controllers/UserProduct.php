<?php

namespace App\Http\Controllers;

use \App\Models\product;

class UserProduct extends Controller
{
    public function index()
    {
            $products = product::all();
            $zhora = "@ZHORA CORPARETED";
            
            return view("index ", compact("products", "zhora"));
    }
}