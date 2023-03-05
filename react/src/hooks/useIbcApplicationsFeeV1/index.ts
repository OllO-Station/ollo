/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useIbcApplicationsFeeV1() {
  const client = useClient();
  const QueryIncentivizedPackets = (query: any, options: any) => {
    const key = { type: 'QueryIncentivizedPackets', query };    
    return useQuery([key], () => {
      const {query } = key
      return  client.IbcApplicationsFeeV1.query.queryIncentivizedPackets(query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryIncentivizedPacket = (packet_id.channel_id: string, packet_id.port_id: string, packet_id.sequence: string,  options: any) => {
    const key = { type: 'QueryIncentivizedPacket',  packet_id.channel_id,  packet_id.port_id,  packet_id.sequence };    
    return useQuery([key], () => {
      const { packet_id.channel_id,  packet_id.port_id,  packet_id.sequence } = key
      return  client.IbcApplicationsFeeV1.query.queryIncentivizedPacket(packet_id.channel_id, packet_id.port_id, packet_id.sequence).then( res => res.data );
    }, options);
  }
  
  const QueryIncentivizedPacketsForChannel = (channel_id: string, port_id: string, query: any, options: any) => {
    const key = { type: 'QueryIncentivizedPacketsForChannel',  channel_id,  port_id, query };    
    return useQuery([key], () => {
      const { channel_id,  port_id,query } = key
      return  client.IbcApplicationsFeeV1.query.queryIncentivizedPacketsForChannel(channel_id, port_id, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryTotalRecvFees = (packet_id.channel_id: string, packet_id.port_id: string, packet_id.sequence: string,  options: any) => {
    const key = { type: 'QueryTotalRecvFees',  packet_id.channel_id,  packet_id.port_id,  packet_id.sequence };    
    return useQuery([key], () => {
      const { packet_id.channel_id,  packet_id.port_id,  packet_id.sequence } = key
      return  client.IbcApplicationsFeeV1.query.queryTotalRecvFees(packet_id.channel_id, packet_id.port_id, packet_id.sequence).then( res => res.data );
    }, options);
  }
  
  const QueryTotalAckFees = (packet_id.channel_id: string, packet_id.port_id: string, packet_id.sequence: string,  options: any) => {
    const key = { type: 'QueryTotalAckFees',  packet_id.channel_id,  packet_id.port_id,  packet_id.sequence };    
    return useQuery([key], () => {
      const { packet_id.channel_id,  packet_id.port_id,  packet_id.sequence } = key
      return  client.IbcApplicationsFeeV1.query.queryTotalAckFees(packet_id.channel_id, packet_id.port_id, packet_id.sequence).then( res => res.data );
    }, options);
  }
  
  const QueryTotalTimeoutFees = (packet_id.channel_id: string, packet_id.port_id: string, packet_id.sequence: string,  options: any) => {
    const key = { type: 'QueryTotalTimeoutFees',  packet_id.channel_id,  packet_id.port_id,  packet_id.sequence };    
    return useQuery([key], () => {
      const { packet_id.channel_id,  packet_id.port_id,  packet_id.sequence } = key
      return  client.IbcApplicationsFeeV1.query.queryTotalTimeoutFees(packet_id.channel_id, packet_id.port_id, packet_id.sequence).then( res => res.data );
    }, options);
  }
  
  const QueryPayee = (channel_id: string, relayer: string,  options: any) => {
    const key = { type: 'QueryPayee',  channel_id,  relayer };    
    return useQuery([key], () => {
      const { channel_id,  relayer } = key
      return  client.IbcApplicationsFeeV1.query.queryPayee(channel_id, relayer).then( res => res.data );
    }, options);
  }
  
  const QueryCounterpartyPayee = (channel_id: string, relayer: string,  options: any) => {
    const key = { type: 'QueryCounterpartyPayee',  channel_id,  relayer };    
    return useQuery([key], () => {
      const { channel_id,  relayer } = key
      return  client.IbcApplicationsFeeV1.query.queryCounterpartyPayee(channel_id, relayer).then( res => res.data );
    }, options);
  }
  
  const QueryFeeEnabledChannels = (query: any, options: any) => {
    const key = { type: 'QueryFeeEnabledChannels', query };    
    return useQuery([key], () => {
      const {query } = key
      return  client.IbcApplicationsFeeV1.query.queryFeeEnabledChannels(query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryFeeEnabledChannel = (channel_id: string, port_id: string,  options: any) => {
    const key = { type: 'QueryFeeEnabledChannel',  channel_id,  port_id };    
    return useQuery([key], () => {
      const { channel_id,  port_id } = key
      return  client.IbcApplicationsFeeV1.query.queryFeeEnabledChannel(channel_id, port_id).then( res => res.data );
    }, options);
  }
  
  return {QueryIncentivizedPackets,QueryIncentivizedPacket,QueryIncentivizedPacketsForChannel,QueryTotalRecvFees,QueryTotalAckFees,QueryTotalTimeoutFees,QueryPayee,QueryCounterpartyPayee,QueryFeeEnabledChannels,QueryFeeEnabledChannel,
  }
}