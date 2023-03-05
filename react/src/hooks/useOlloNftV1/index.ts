/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloNftV1() {
  const client = useClient();
  const QuerySupply = (denom_id: string, query: any, options: any) => {
    const key = { type: 'QuerySupply',  denom_id, query };    
    return useQuery([key], () => {
      const { denom_id,query } = key
      return  client.OlloNftV1.query.querySupply(denom_id, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryNFTsOfOwner = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryNFTsOfOwner', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloNftV1.query.queryNFTsOfOwner(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryCollection = (denom_id: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryCollection',  denom_id, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { denom_id,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloNftV1.query.queryCollection(denom_id, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryDenom = (denom_id: string,  options: any) => {
    const key = { type: 'QueryDenom',  denom_id };    
    return useQuery([key], () => {
      const { denom_id } = key
      return  client.OlloNftV1.query.queryDenom(denom_id).then( res => res.data );
    }, options);
  }
  
  const QueryDenoms = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryDenoms', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloNftV1.query.queryDenoms(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryNFT = (denom_id: string, token_id: string,  options: any) => {
    const key = { type: 'QueryNFT',  denom_id,  token_id };    
    return useQuery([key], () => {
      const { denom_id,  token_id } = key
      return  client.OlloNftV1.query.queryNFT(denom_id, token_id).then( res => res.data );
    }, options);
  }
  
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloNftV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  return {QuerySupply,QueryNFTsOfOwner,QueryCollection,QueryDenom,QueryDenoms,QueryNFT,QueryParams,
  }
}