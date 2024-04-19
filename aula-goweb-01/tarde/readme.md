# Reference

I`ve use the following queries to test:

(GET ALL)
http://localhost:8080/transactions

(GET ONE)
http://localhost:8080/transactions/1
http://localhost:8080/transactions/2

(GET FILTERED)

CASE: ID = [1, 2, 3] OR Currency = [USD] OR Amount = [10000.00]
http://localhost:8080/transactions?ID=1,2,3&Currency=USD&Amount=10000.00

CASE: ID = [1, 2, 3] OR Currency = [USD, BRL] OR Amount = [10000.00]
http://localhost:8080/transactions?ID=1,2,3&Currency=USD,BRL&Amount=10000.00

CASE: ID = [1] OR Currency = [BRL]
http://localhost:8080/transactions?ID=1&Currency=BRL

CASE: ID = [1] OR Amount = [10000.00] OR Currency = [BRL]
http://localhost:8080/transactions?ID=1&Amount=10000.00&Currency=BRL

CASE: ID=[1,2] OR Amount = [10000.00] OR Currency = [BRL, USD] OR Date = [2020-01-12]
http://localhost:8080/transactions?ID=1,2&Amount=10000.00&Currency=BRL,USD&Date=2020-01-02