#!/bin/bash

DESC_NOTFOUND=`echo -e "ERROR:\n  Code: NotFound\n  Message: company not found\n"`
DESC_INVALID=`echo -e "ERROR:\n  Code: InvalidArgument\n  Message: invalid DescribeCompanyV1Request.CompanyId: value must be greater than 0"`

echo
ANS=`grpcurl -d '{"company_id": 1}' -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/DescribeCompanyV1 2>&1`
echo 'grpcurl SescriveCompanyV1 id:1'
if [[ "$ANS" = "$DESC_NOTFOUND" ]]
then
    echo "Ok";
else
    echo "Error"
fi

echo
ANS=`grpcurl -d '{"company_id": 0}' -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/DescribeCompanyV1 2>&1`
echo 'grpcurl SescriveCompanyV1 id:0'
if [[ "$ANS" = "$DESC_INVALID" ]]
then
    echo "Ok";
else
    echo "Error"
fi