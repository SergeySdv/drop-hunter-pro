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

export class FlowService {
  /**
   *
   */
  flowServiceCreateFlow(
    params: {
      /**  */
      body: CreateFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/flow/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceDeleteFlow(
    params: {
      /**  */
      body: DeleteFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DeleteFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/flow/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceGetFlow(
    params: {
      /**  */
      body: GetFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/flow/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceListFlow(
    params: {
      /**  */
      body: ListFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/flow/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface CreateFlowRequest {
  /**  */
  label: string;

  /**  */
  tasks: Task[];
}

export interface CreateFlowResponse {
  /**  */
  flow: flow_Flow;
}

export interface DelayTask {
  /**  */
  duration: string;

  /**  */
  waitFor?: Date;

  /**  */
  random: boolean;

  /**  */
  minRandom?: string;

  /**  */
  maxRandom?: string;
}

export interface DeleteFlowRequest {
  /**  */
  id: string;
}

export interface DeleteFlowResponse {}

export interface GetFlowRequest {
  /**  */
  id: string;
}

export interface GetFlowResponse {
  /**  */
  flow: flow_Flow;
}

export interface ListFlowRequest {}

export interface ListFlowResponse {
  /**  */
  flows: flow_Flow[];
}

export interface MockTask {}

export interface StargateBridgeTask {
  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;
}

export interface Task {
  /**  */
  weight: string;

  /**  */
  taskType: TaskType;

  /**  */
  description: string;

  /**  */
  stargateBridgeTask?: StargateBridgeTask;

  /**  */
  mockTask?: MockTask;

  /**  */
  delayTask?: DelayTask;

  /**  */
  withdrawExchangeTask?: WithdrawExchangeTask;
}

export interface WithdrawExchangeTask {
  /**  */
  withdrawerId: string;

  /**  */
  network: string;

  /**  */
  amountMin: string;

  /**  */
  amountMax: string;

  /**  */
  token: string;

  /**  */
  withdrawOrderId: string;

  /**  */
  amount?: string;

  /**  */
  txId?: string;
}

export interface flow_Flow {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  tasks: Task[];

  /**  */
  nextId?: string;

  /**  */
  createdAt: Date;
}

export enum Network {
  'ARBITRUM' = 'ARBITRUM',
  'OPTIMISM' = 'OPTIMISM',
  'BinanaceBNB' = 'BinanaceBNB',
  'Etherium' = 'Etherium',
  'POLIGON' = 'POLIGON',
  'AVALANCHE' = 'AVALANCHE'
}

export enum TaskType {
  'StargateBridge' = 'StargateBridge',
  'Mock' = 'Mock',
  'Delay' = 'Delay',
  'WithdrawExchange' = 'WithdrawExchange'
}

export enum Token {
  'USDT' = 'USDT',
  'ETH' = 'ETH',
  'USDC' = 'USDC',
  'STG' = 'STG',
  'BNB' = 'BNB',
  'MATIC' = 'MATIC',
  'AVAX' = 'AVAX'
}
