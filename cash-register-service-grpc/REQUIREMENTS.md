# Cash Register Receipt Printer Service
The goal of this challenge is to create a web service for generating receipts for
a typical supermarket checkout register.  This document describes the behavior
of the service and its inputs and outputs.

## Problem Description
The web service which you need to write (which we will refer to as
"cash register") shall be able to store a catalog of products and prices,  
as well as receive purchase transactions containing a list of items from the catalog
scanned by the cash register's optical scanner. Your job is to write a web service 
that takes these transactions and, using the data from the catalog, returns an itemized 
receipt for each customer. You may choose any service style you would like such as 
REST, gRPC, GraphQL, etc..

## Establishing the Catalog
Create a service API that adds items to a store catalog.
There are 2 pieces of data for each item in the catalog:
1. Item Name
2. Item Unit Price

You can store the catalog in a data store of your choice. 

## Purchase Transaction
Create a service API that receives a list of scanned items from a scanner and returns a receipt for the transaction.
### Transaction Request
Every item is scanned individually. If a customer purchases two apples, the apple line
item should occur twice in the request.

### Transaction Response
For each transaction request, your service should return a customer's receipt.   
A receipt should contain the following data per kind of item purchased. 
 * the quantity of the item purchased
 * the item's name
 * the items' unit price
 * the "extended" price 

The "extended price" is computed by multiplying the quantity by the customer's price 
for that item.

A receipt should also contain the TOTAL for the entire transaction.
 
Note that items should appear only once on a receipt.  In particular, multiple instances 
of the same item should be consolidated on one line item in the response no matter 
where they appeared in the transaction file.

Items should be printed in the order they were FIRST encountered in the transaction request. 



## Evaluation:
Please write tests  that will demonstrate functionality and correctness of your web service.

Please be prepared to discuss the design choices you make in api and data modeling as well as storage.
