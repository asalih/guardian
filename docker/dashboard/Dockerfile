FROM mcr.microsoft.com/dotnet/core/sdk:3.0

RUN apt-get update
RUN apt-get install -y git

RUN git clone https://github.com/asalih/GuardianUI

WORKDIR ./GuardianUI

ENV ASPNETCORE_ENVIRONMENT=Production

RUN dotnet build --configuration Release
WORKDIR ./Guardian.Web.UI/bin/Release/netcoreapp3.0/

RUN sed -i 's,POSTGRES_PASSWORD,7eb12540045a4bc0b474a7efb91ebd39,g' appsettings.json && \
    sed -i 's,POSTGRES_HOST,db,g' appsettings.json && \
    sed -i 's,POSTGRES_USER,guardian,g' appsettings.json && \
    sed -i 's,POSTGRES_DB,guardiandb,g' appsettings.json

CMD ["dotnet", "Guardian.Web.UI.dll"]

EXPOSE 8080