package exone

/* 
    We receive notification via a webhook, for new candidates status change. The webhook contains only a little part of the resource we need in our platfomr. 
    For each request we need to retrieve the relevant infomration needed and store them in a database.
    The operation of retrieval is syncronous.    
    Lately we noticed an increased number of errors with http status 429.
    You are assigned to the team that needs to undestand and fix the problem

*/
