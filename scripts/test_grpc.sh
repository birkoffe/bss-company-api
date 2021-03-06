#!/bin/bash

NOTFOUND=`echo -e "ERROR:\n  Code: NotFound\n  Message: company not found\n"`
INVALID=`echo -e "ERROR:\n  Code: InvalidArgument\n  Message: invalid DescribeCompanyV1Request.CompanyId: value must be greater than 0"`
NOTIMP=`echo -e "ERROR:\n  Code: Internal\n  Message: not implemented"`

echo
ANS=`grpcurl -d '{"company_id": 1}' -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/DescribeCompanyV1 2>&1`
echo 'grpcurl SescriveCompanyV1 id:1'
if [[ "$ANS" = "$NOTFOUND" ]]
then
    echo "Ok";
else
    echo "Error"
fi

echo
ANS=`grpcurl -d '{"company_id": 0}' -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/DescribeCompanyV1 2>&1`
echo 'grpcurl SescriveCompanyV1 id:0'
if [[ "$ANS" = "$INVALID" ]]
then
    echo "Ok";
else
    echo "Error"
fi

echo 
ANS=`grpcurl -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/CreateCompanyV1 2>&1`
echo 'grpcurl SescriveCompanyV1'
if [[ "$ANS" = "$NOTIMP" ]]
then
    echo "Ok";
else
    echo "Error"
fi

echo 
ANS=`grpcurl -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/ListCompanyV1 2>&1`
echo 'grpcurl ListCompanyV1'
if [[ "$ANS" = "$NOTIMP" ]]
then
    echo "Ok";
else
    echo "Error"
fi


echo 
ANS=`grpcurl -plaintext localhost:8082  ozonmp.bss_company_api.v1.BssCompanyApiService/RemoveCompanyV1 2>&1`
echo 'grpcurl RemoveCompanyV1'
if [[ "$ANS" = "$NOTIMP" ]]
then
    echo "Ok";
else
    echo "Error"
fi