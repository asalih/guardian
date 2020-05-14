FROM mcr.microsoft.com/dotnet/core/sdk:3.0

RUN apt-get update
RUN apt-get install -y git
RUN dotnet tool install -g Cake.Tool && \
    echo 'export PATH=$PATH:/root/.dotnet/tools' >> $BASH_ENV && \
    source $BASH_ENV

RUN git clone https://github.com/asalih/GuardianUI

WORKDIR ./GuardianUI

RUN dotnet cake build.cake --runtime=linux-x64

WORKDIR ./publish

RUN sed -i 's,POSTGRES_PASSWORD,7eb12540045a4bc0b474a7efb91ebd39,g' appsettings.json && \
    sed -i 's,POSTGRES_HOST,db,g' appsettings.json && \
    sed -i 's,POSTGRES_USER,guardian,g' appsettings.json && \
    sed -i 's,POSTGRES_DB,guardiandb,g' appsettings.json

CMD ["dotnet", "Guardian.Web.UI.dll"]

EXPOSE 8080