/** Generate by swagger-axios-codegen */
// @ts-nocheck
/* eslint-disable */

/** Generate by swagger-axios-codegen */
// @ts-nocheck
/* eslint-disable */
export interface IRequestOptions {
  headers?: any;
  /** only in axios interceptor config*/
  loading: boolean;
  showError: boolean;
}

export interface IRequestPromise<T = any> extends Promise<IRequestResponse<T>> {}

export interface IRequestResponse<T = any> {
  data: T;
  status: number;
  statusText: string;
  headers: any;
  config: any;
  request?: any;
}

export interface IRequestInstance {
  (config: any): IRequestPromise;
  (url: string, config?: any): IRequestPromise;
  request<T = any>(config: any): IRequestPromise<T>;
}

export interface IRequestConfig {
  method?: any;
  headers?: any;
  url?: any;
  data?: any;
  params?: any;
}

// Add options interface
export interface ServiceOptions {
  axios?: IRequestInstance;
  /** only in axios interceptor config*/
  loading: boolean;
  showError: boolean;
}

// Add default options
export const serviceOptions: ServiceOptions = {};

// Instance selector
export function axios(configs: IRequestConfig, resolve: (p: any) => void, reject: (p: any) => void): Promise<any> {
  if (serviceOptions.axios) {
    return serviceOptions.axios
      .request(configs)
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  } else {
    throw new Error('please inject yourself instance like axios  ');
  }
}

export function getConfigs(method: string, contentType: string, url: string, options: any): IRequestConfig {
  const configs: IRequestConfig = {
    loading: serviceOptions.loading,
    showError: serviceOptions.loading,
    ...options,
    method,
    url
  };
  configs.headers = {
    ...options.headers,
    'Content-Type': contentType
  };
  return configs;
}

export const basePath = '';

export interface IList<T> extends Array<T> {}
export interface List<T> extends Array<T> {}
export interface IDictionary<TValue> {
  [key: string]: TValue;
}
export interface Dictionary<TValue> extends IDictionary<TValue> {}

export interface IListResult<T> {
  items?: T[];
}

export class ListResultDto<T> implements IListResult<T> {
  items?: T[];
}

export interface IPagedResult<T> extends IListResult<T> {
  totalCount?: number;
  items?: T[];
}

export class PagedResultDto<T = any> implements IPagedResult<T> {
  totalCount?: number;
  items?: T[];
}

// customer definition
// empty

export class HelperService {
  /**
   *
   */
  helperServiceGetBillingHistory(
    params: {
      /**  */
      body: GetBillingHistoryReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetBillingHistoryRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/billing/history';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceCreateOrder(
    params: {
      /**  */
      body: CreateOrderReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateOrderRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/billing/order/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceGetOrderHistory(
    params: {
      /**  */
      body: GetOrderHistoryReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetOrderHistoryRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/billing/order/history';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceGetOrderStatus(
    params: {
      /**  */
      body: GetOrderStatusReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetOrderStatusRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/billing/order/status';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceCastWei(
    params: {
      /**  */
      body: CastWEIRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CastWEIResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/helper/cast-wei';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceEstimateStargateBridgeFee(
    params: {
      /**  */
      body: EstimateStargateBridgeFeeRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<EstimateStargateBridgeFeeResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/helper/stargate-bridge-fee/estimate';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceValidatePk(
    params: {
      /**  */
      body: ValidatePKRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ValidatePKResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/helper/validatePK';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceValidateProxy(
    params: {
      /**  */
      body: ValidateProxyRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ValidateProxyResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/helper/validateProxy';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceTransactionsDailyImpact(
    params: {
      /**  */
      body: TransactionsDailyImpactReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<TransactionsDailyImpactRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/impact/daily';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  helperServiceGetUser(
    params: {
      /**  */
      body: GetUserRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetUserResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/user';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface AmUni {
  /**  */
  gwei?: string;

  /**  */
  eth?: string;

  /**  */
  usd?: string;

  /**  */
  network?: Network;

  /**  */
  wei?: string;
}

export interface CastWEIRequest {
  /**  */
  wei: string;

  /**  */
  network: Network;
}

export interface CastWEIResponse {
  /**  */
  am: AmUni;
}

export interface CreateOrderReq {
  /**  */
  userId?: string;

  /**  */
  am: string;
}

export interface CreateOrderRes {
  /**  */
  id: string;

  /**  */
  coinAddrUrl: string;

  /**  */
  am: number;

  /**  */
  toWallet: string;
}

export interface EstimateStargateBridgeFeeRequest {
  /**  */
  from: Network;

  /**  */
  to: Network;

  /**  */
  profileId: string;
}

export interface EstimateStargateBridgeFeeResponse {
  /**  */
  wei: string;

  /**  */
  usd: string;

  /**  */
  eth: string;

  /**  */
  error?: string;
}

export interface GetBillingHistoryReq {
  /**  */
  offset: string;
}

export interface GetBillingHistoryRes {
  /**  */
  records: TaskHistoryRecord[];
}

export interface GetOrderHistoryReq {}

export interface GetOrderHistoryRes {
  /**  */
  orders: Order[];
}

export interface GetOrderStatusReq {
  /**  */
  orderId: string;
}

export interface GetOrderStatusRes {
  /**  */
  status: string;
}

export interface GetUserRequest {}

export interface GetUserResponse {
  /**  */
  id: string;

  /**  */
  email: string;

  /**  */
  funds: string;

  /**  */
  taskPrice: string;

  /**  */
  payableTasks: TaskType[];

  /**  */
  nonpayableTasks: TaskType[];
}

export interface Order {
  /**  */
  id: string;

  /**  */
  net: string;

  /**  */
  coinAddrUrl: string;

  /**  */
  status: string;

  /**  */
  createdAt: string;

  /**  */
  confirmedAt: string;

  /**  */
  am: number;

  /**  */
  toWallet: string;
}

export interface TaskHistoryRecord {
  /**  */
  processId: string;

  /**  */
  taskId: string;

  /**  */
  taskType: string;

  /**  */
  taskPrice: number;
}

export interface TransactionsDailyImpactReq {}

export interface TransactionsDailyImpactRes {
  /**  */
  myImpact: string;

  /**  */
  totalImpact: string;

  /**  */
  topImpact: string;
}

export interface ValidatePKRequest {
  /**  */
  pk: string;
}

export interface ValidatePKResponse {
  /**  */
  valid: boolean;

  /**  */
  walletId?: string;
}

export interface ValidateProxyRequest {
  /**  */
  proxy: string;
}

export interface ValidateProxyResponse {
  /**  */
  valid: boolean;

  /**  */
  errorMessage: string;

  /**  */
  countryName: string;

  /**  */
  countryCode: string;

  /**  */
  ip: string;
}

export enum Network {
  'ARBITRUM' = 'ARBITRUM',
  'OPTIMISM' = 'OPTIMISM',
  'BinanaceBNB' = 'BinanaceBNB',
  'Etherium' = 'Etherium',
  'POLIGON' = 'POLIGON',
  'AVALANCHE' = 'AVALANCHE',
  'GOERLIETH' = 'GOERLIETH',
  'ZKSYNCERA' = 'ZKSYNCERA',
  'ZKSYNCERATESTNET' = 'ZKSYNCERATESTNET',
  'ZKSYNCLITE' = 'ZKSYNCLITE'
}

export enum TaskType {
  'StargateBridge' = 'StargateBridge',
  'Mock' = 'Mock',
  'Delay' = 'Delay',
  'WithdrawExchange' = 'WithdrawExchange',
  'OkexDeposit' = 'OkexDeposit',
  'TestNetBridgeSwap' = 'TestNetBridgeSwap',
  'SnapshotVote' = 'SnapshotVote',
  'OkexBinance' = 'OkexBinance',
  'Swap1inch' = 'Swap1inch',
  'SyncSwap' = 'SyncSwap',
  'ZkSyncOfficialBridgeToEthereum' = 'ZkSyncOfficialBridgeToEthereum',
  'OrbiterBridge' = 'OrbiterBridge',
  'ZkSyncOfficialBridgeFromEthereum' = 'ZkSyncOfficialBridgeFromEthereum',
  'WETH' = 'WETH',
  'MuteioSwap' = 'MuteioSwap',
  'SyncSwapLP' = 'SyncSwapLP',
  'MaverickSwap' = 'MaverickSwap',
  'SpaceFISwap' = 'SpaceFISwap',
  'VelocoreSwap' = 'VelocoreSwap',
  'IzumiSwap' = 'IzumiSwap',
  'VeSyncSwap' = 'VeSyncSwap',
  'EzkaliburSwap' = 'EzkaliburSwap',
  'ZkSwap' = 'ZkSwap'
}
