<?php

use App\Http\Controllers\UserProduct;
use Illuminate\Support\Facades\Route;

Route::get("/", [UserProduct::class, "index"]);