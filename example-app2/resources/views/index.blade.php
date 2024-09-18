<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> Lara Croft</title>
</head>
<body>
    <style>
        .cards {
            display: flex;
            justify-content: center;
            flex-wrap: wrap;
            margin: 20px;
            flex-direction: row;
        }
    
        .card {
            background-color: #f9f9f9;
            border: 1px solid #ddd;
            border-radius: 10px;
            padding: 15px;
            margin: 10px;
            text-align: center;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
            width: 200px;
            transition: transform 0.2s;
        }
    
        .card:hover {
            transform: scale(1.05);
        }
    
        .name {
            font-weight: bold;
            font-size: 1.2em;
        }
    
        .amount, .cost {
            font-size: 1em;
        }
    
        button {
            background-color: #28a745;
            color: white;
            border: none;
            border-radius: 5px;
            padding: 10px 15px;
            cursor: pointer;
            transition: background-color 0.2s;
        }
    
        button:hover {
            background-color: #218838;
        }
    </style>
    <div class="cards">  
    @foreach ($products as $product)
        
        <div class="card">
            <p class="name">{{ $product->name}}</p>
            <p class="amount">{{ $product->amount }}</p>
            <p class="cost">{{ $product->cost }}</p>
            <button>Купить</button>
        </div>
    
    @endforeach

</div>
    
   
</body>
<footer>{{ $zhora}}</footer>
</html>