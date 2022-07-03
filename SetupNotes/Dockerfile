// Docker File

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y wget git g++ unixodbc-dev unixodbc build-essential && \
    apt-get install -y coinor-cbc glpk-utils libglpk-dev glpk-doc python3-swiglpk libpq-dev

RUN pip install --upgrade pip

WORKDIR /code

COPY pip.conf ./pip.conf
COPY requirements.txt ./requirements.txt

ENV PIP_CONFIG_FILE=./pip.conf
RUN pip install --trusted-host pypi.org --trusted-host files.pythonhosted.org -r requirements.txt

COPY optopy_grpc ./optopy_grpc

EXPOSE 5050

ENV PYTHONPATH "${PYTHONPATH}:/code"

CMD ["python", "optopy_grpc/server.py"]
