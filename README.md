# Capital-Adequacy
Capital-adequacy command is used to capital-adequacy create table on basis of 3 inputs and send  metrics over AWS cloudwatch such as ValueAtRisk, TimeTaken by the job, PositionAge, SymbolAge.
* Sum of all realisedPNL in float which is converted into money value.
* Get the latest entry from PositionSnapshot table.
* Get the latest entry from CashRequirement table.
When Capital-Adequacy command runs, all these values get accumulated and get converted from USD to AUD.
### To run this capital-adequacy command
* Go inside airflow-report>>capital-adequacy
* Type docker-compose up
* By running this docker-compose up it will automatically download all the required images, such as mysql, go, all the dependecies.	
### Configuration of aws for cloudwatch
* Before sending custom metrics to cloudwatch, first we have to create ec2 instance in aws:
for creating ec2 instance we have to set IAM policies for the service cloudwatch and set action as PutMetricData, give policy name and policy description. (after setting up click on create policy).

* Now create Role of type EC2, click on permission and search for policy which you have created, then click on create role.

* Now we will launch ec2 instance, select amazone one linux, in configure instance details select IAM role(which you have created) and rest setting will be default, storage will be default, give tag name which is key and value (which will identify the instance), select the existing security group of port 22, then click on review and launch with the key pair (which you have declared) then click on launch instance.

* After creating ec2 instance, run this command into terminal, to check wheather the instance is created or not.
* ssh ec2-user@IPv4 Public IP -i ".pem file" (here IPv4 Public IP is the ip address which can be found at ec2 instance consol, .pem file is hte file which is given by aws)
* If permission is denied give permission in terminal by typing command 400 chmod
* After getting into ec2 instance set aws credentials by typing command "aws configure" into ec2 terminal, set the access key id, secret access key, region and output format.
* Type nano /etc/hosts into this file set "Private IPs monitoring.us-east-2c.amazonaws.com", here us-east-2c is used as region, we can configure according to the need, Private IPs we will find at ec2 consol.
* When code is executed successfully, to see metrics over cloudwatch, login to aws, go to cloudwatch, in cloudwatch panel go to Metrics and here we can see all the custom metrics which are send from the program.
	
