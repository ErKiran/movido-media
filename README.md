# movido-media

## To run the code 

1. Create a environment file from .env.example rename it to .env and fill it with necessary variables for database and mail server mailtrap for an instance. 

2. Run the migration script first 
```golang
 go run migrations/migrate.go
 ```

3. Change the cron job timer to run it instantly.

### Database 

For the simplicty I choose to use three table customers, products and contracts. 

I have added a initial dataseed to get started.

I have skipped updating the database after the invoice has been generated and sent for customer to make things simpler for this poc. 

Inorder to handle the case of error in any of the pipeline process I have use exponential retry backoff strategy. It will try to redo the process 5 time with the interval of 5 minutes. Also, I have utilized the logging to records error if any.

### PDF 
I have used a disk storage instead of sending the pdf to cloud for the simplicity. 

### Billing Logic 

There are many strategy for the billing lifecycle. In this implementation, billing date is kept as it is. For example if you have subscribed for the service in 15th of this months it will be renewed in 15th of next month.