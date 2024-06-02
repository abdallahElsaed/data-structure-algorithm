# Hash Table Data Structure

## What is hashing?
* It has many names:
    * hash table
    * dictionary
    * map
    * associative array 
* It has items all those items has own key. [ key : value ]
* each key is translated by a hash function into a distinct index in an array
* you can make those operation on it:
    1. Insert
    2. Delete
    3. Search => you can search by key to get value.
    > All time those operation toke is constant time O(1)

## Why don't store value in array direct and use his index to be a key, What is problem?
> Why don't get the ASCII code for string kay or any type amd store on it array?
1. Because the key may be not integer.
2. gigantic memory hog. (we will have a big array)
### Solve first problem 
* we do some thing called 'Prehash, this way take the string or any type and convert to integer (bits)
* ex: string='a' => hash('a') = 65  
### solve second problem
* after convert to integer the position of the key be a gigantic.
* so you can do some thing to make the key equal array size, this is called hash function, you can passing the key to this function and return number equal or less than size or this array.