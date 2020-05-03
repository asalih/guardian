# Guardian Web Application Firewall
Guardian is the open source web application firewall based on ModSecurity SecRules format. 

# How it works!
Guardian locates in front of your web server and if incoming traffic valid then the Guardian passes it to the target server.

![Diagram](images/guardian.png)

## Guardian Nameserver
[Guardian Nameserver](https://github.com/asalih/guardian_ns) To route web traffic through the Guardian, update the nameservers at your domain registrar to resolve your domain’s DNS with Guardian's nameservers.

## Guardian Dashboard
[Guardian Dashboard](https://github.com/asalih/GuardianUI) To managing your rules and domains.