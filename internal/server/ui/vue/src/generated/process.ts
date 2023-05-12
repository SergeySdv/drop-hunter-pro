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

export class ProcessService {
  /**
   *
   */
  processServiceCancelProcess(
    params: {
      /**  */
      body: CancelProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CancelProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/cancel';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceCreateProcess(
    params: {
      /**  */
      body: CreateProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceGetProcess(
    params: {
      /**  */
      body: GetProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceListProcess(
    params: {
      /**  */
      body: ListProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceResumeProcess(
    params: {
      /**  */
      body: ResumeProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ResumeProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/resume';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceRetryProcess(
    params: {
      /**  */
      body: RetryProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<RetryProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/retry';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceStopProcess(
    params: {
      /**  */
      body: StopProcessRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<StopProcessResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/stop';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceGetProcessTaskHistory(
    params: {
      /**  */
      body: GetProcessTaskHistoryRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetProcessTaskHistoryResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/task/history';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceGetProcessUpdatedAt(
    params: {
      /**  */
      body: GetProcessUpdatedAtRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetProcessUpdatedAtResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/process/updatedAt';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface CancelProcessRequest {
  /**  */
  processId: string;
}

export interface CancelProcessResponse {}

export interface CreateProcessRequest {
  /**  */
  flowId: string;

  /**  */
  profileIds: string[];
}

export interface CreateProcessResponse {
  /**  */
  process: Process;
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

export interface GetProcessRequest {
  /**  */
  id: string;
}

export interface GetProcessResponse {
  /**  */
  process: Process;
}

export interface GetProcessTaskHistoryRequest {
  /**  */
  taskId: string;
}

export interface GetProcessTaskHistoryResponse {
  /**  */
  records: ProcessTaskHistoryRecord[];
}

export interface GetProcessUpdatedAtRequest {
  /**  */
  processId: string;
}

export interface GetProcessUpdatedAtResponse {
  /**  */
  updatedAt: Date;
}

export interface ListProcessRequest {}

export interface ListProcessResponse {
  /**  */
  processes: Process[];
}

export interface MockTask {}

export interface Process {
  /**  */
  id: string;

  /**  */
  status: ProcessStatus;

  /**  */
  profiles: ProcessProfile[];

  /**  */
  flowId: string;

  /**  */
  createdAt: Date;

  /**  */
  updatedAt: Date;

  /**  */
  finishedAt?: Date;

  /**  */
  startedAt?: Date;

  /**  */
  flowLabel: string;
}

export interface ProcessProfile {
  /**  */
  profileId: string;

  /**  */
  weight: string;

  /**  */
  tasks: ProcessTask[];

  /**  */
  status: ProcessStatus;

  /**  */
  id: string;
}

export interface ProcessTask {
  /**  */
  task: Task;

  /**  */
  status: ProcessStatus;

  /**  */
  transactions: string[];

  /**  */
  finishedAt?: Date;

  /**  */
  startedAt?: Date;

  /**  */
  error?: string;

  /**  */
  id: string;
}

export interface ProcessTaskHistoryRecord {
  /**  */
  id: string;

  /**  */
  taskId: string;

  /**  */
  startedAt: Date;

  /**  */
  finishedAt?: Date;

  /**  */
  startStatus: ProcessStatus;

  /**  */
  finishStatus?: ProcessStatus;

  /**  */
  msg?: string;
}

export interface ResumeProcessRequest {
  /**  */
  processId: string;
}

export interface ResumeProcessResponse {}

export interface RetryProcessRequest {
  /**  */
  processId: string;

  /**  */
  profileId: string;

  /**  */
  taskId: string;
}

export interface RetryProcessResponse {}

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

export interface StopProcessRequest {
  /**  */
  processId: string;
}

export interface StopProcessResponse {}

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

export enum Network {
  'ARBITRUM' = 'ARBITRUM',
  'OPTIMISM' = 'OPTIMISM',
  'BinanaceBNB' = 'BinanaceBNB',
  'Etherium' = 'Etherium',
  'POLIGON' = 'POLIGON',
  'AVALANCHE' = 'AVALANCHE'
}

export enum ProcessStatus {
  'StatusReady' = 'StatusReady',
  'StatusRunning' = 'StatusRunning',
  'StatusError' = 'StatusError',
  'StatusDone' = 'StatusDone',
  'StatusStop' = 'StatusStop',
  'StatusRetry' = 'StatusRetry'
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
