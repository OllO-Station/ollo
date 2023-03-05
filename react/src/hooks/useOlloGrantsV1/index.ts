/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloGrantsV1() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloGrantsV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryAuctions = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryAuctions', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloGrantsV1.query.queryAuctions(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryAuction = (auction_id: string,  options: any) => {
    const key = { type: 'QueryAuction',  auction_id };    
    return useQuery([key], () => {
      const { auction_id } = key
      return  client.OlloGrantsV1.query.queryAuction(auction_id).then( res => res.data );
    }, options);
  }
  
  const QueryAllowedBidder = (auction_id: string, bidder: string,  options: any) => {
    const key = { type: 'QueryAllowedBidder',  auction_id,  bidder };    
    return useQuery([key], () => {
      const { auction_id,  bidder } = key
      return  client.OlloGrantsV1.query.queryAllowedBidder(auction_id, bidder).then( res => res.data );
    }, options);
  }
  
  const QueryAllowedBidders = (auction_id: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryAllowedBidders',  auction_id, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { auction_id,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloGrantsV1.query.queryAllowedBidders(auction_id, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryBids = (auction_id: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryBids',  auction_id, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { auction_id,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloGrantsV1.query.queryBids(auction_id, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryBid = (auction_id: string, bid_id: string,  options: any) => {
    const key = { type: 'QueryBid',  auction_id,  bid_id };    
    return useQuery([key], () => {
      const { auction_id,  bid_id } = key
      return  client.OlloGrantsV1.query.queryBid(auction_id, bid_id).then( res => res.data );
    }, options);
  }
  
  const QueryVestings = (auction_id: string,  options: any) => {
    const key = { type: 'QueryVestings',  auction_id };    
    return useQuery([key], () => {
      const { auction_id } = key
      return  client.OlloGrantsV1.query.queryVestings(auction_id).then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryAuctions,QueryAuction,QueryAllowedBidder,QueryAllowedBidders,QueryBids,QueryBid,QueryVestings,
  }
}