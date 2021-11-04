import asyncio

from grpclib.client import Channel

from ozonmp.bss_company_api.v1.bss_company_api_grpc import BssCompanyApiServiceStub
from ozonmp.bss_company_api.v1.bss_company_api_pb2 import DescribeCompanyV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = BssCompanyApiServiceStub(channel)

        req = DescribeCompanyV1Request(company_id=1)
        reply = await client.DescribecompanyV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
