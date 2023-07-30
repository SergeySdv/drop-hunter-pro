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
  processServiceDisableAutoRetry(
    params: {
      /**  */
      body: DisableAutoRetryRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DisableAutoRetryResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/auto-retry/disable';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceEnableAutoRetry(
    params: {
      /**  */
      body: EnableAutoRetryRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<EnableAutoRetryResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/auto-retry/enable';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
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
      let url = basePath + '/api/gw/v1/process/cancel';

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
      let url = basePath + '/api/gw/v1/process/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceEstimateCost(
    params: {
      /**  */
      body: EstimateCostRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<EstimateCostResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/estimate-cost';

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
      let url = basePath + '/api/gw/v1/process/get';

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
      let url = basePath + '/api/gw/v1/process/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceGetProfileTransactions(
    params: {
      /**  */
      body: GetProfileTransactionsReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetProfileTransactionsRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/profile/transactions';

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
      let url = basePath + '/api/gw/v1/process/resume';

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
      let url = basePath + '/api/gw/v1/process/retry';

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
      let url = basePath + '/api/gw/v1/process/stop';

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
      let url = basePath + '/api/gw/v1/process/task/history';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceGetTaskSettings(
    params: {
      /**  */
      body: GetTaskSettingsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetTaskSettingsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/task/settings';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceSetTaskSettings(
    params: {
      /**  */
      body: SetTaskSettingsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SetTaskSettingsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/task/settings/set';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceSkipProcessTask(
    params: {
      /**  */
      body: SkipProcessTaskRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SkipProcessTaskResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/task/skip';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  processServiceGetTaskTransactions(
    params: {
      /**  */
      body: GetTaskTransactionsReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetTaskTransactionsRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/process/task/transactions';

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
      let url = basePath + '/api/gw/v1/process/updatedAt';

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

export interface Amount {
  /**  */
  sendAll?: boolean;

  /**  */
  sendPercent?: number;

  /**  */
  sendAmount?: number;

  /**  */
  sendValue?: string;

  /**  */
  send?: AmUni;

  /**  */
  balance?: AmUni;

  /**  */
  gasEstimated?: AmUni;
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

  /**  */
  randomDuration?: string;
}

export interface DisableAutoRetryRequest {
  /**  */
  processId: string;
}

export interface DisableAutoRetryResponse {}

export interface EnableAutoRetryRequest {
  /**  */
  processId: string;
}

export interface EnableAutoRetryResponse {}

export interface EstimateCostRequest {
  /**  */
  processId: string;

  /**  */
  profileId: string;

  /**  */
  taskId: string;
}

export interface EstimateCostResponse {
  /**  */
  error?: string;

  /**  */
  data?: EstimationTx;
}

export interface EstimationTx {
  /**  */
  balance: AmUni;

  /**  */
  value: AmUni;

  /**  */
  gas: AmUni;

  /**  */
  gasLimit: AmUni;

  /**  */
  gasPrice: AmUni;

  /**  */
  gasValuePercent: string;
}

export interface EzkaliburSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
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

export interface GetProfileTransactionsReq {
  /**  */
  profileId: string;
}

export interface GetProfileTransactionsRes {
  /**  */
  transactions: Transaction[];
}

export interface GetTaskSettingsRequest {
  /**  */
  taskType: TaskType;

  /**  */
  network: Network;
}

export interface GetTaskSettingsResponse {
  /**  */
  gasLimit?: AmUni;
}

export interface GetTaskTransactionsReq {
  /**  */
  taskId: string;
}

export interface GetTaskTransactionsRes {
  /**  */
  transactions: Transaction[];
}

export interface IzumiSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface ListProcessRequest {
  /**  */
  statuses: ProcessStatus[];

  /**  */
  offset: string;
}

export interface ListProcessResponse {
  /**  */
  processes: Process[];
}

export interface MaverickSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface MockTask {}

export interface MuteioSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface OkexBinanaceTask {
  /**  */
  okexWithdrawerId: string;

  /**  */
  okexToken: string;

  /**  */
  okexNetwork: string;

  /**  */
  binanaceDepositAddr: string;

  /**  */
  txId?: string;

  /**  */
  withdrawTxId?: string;
}

export interface OkexDepositTask {
  /**  */
  network: Network;

  /**  */
  okexAccName?: string;

  /**  */
  okexAddr?: string;

  /**  */
  token: Token;

  /**  */
  address?: string;

  /**  */
  txId?: string;

  /**  */
  txComplete?: boolean;

  /**  */
  subMainTransfer?: boolean;

  /**  */
  amount?: Amount;

  /**  */
  tx?: TaskTx;
}

export interface OrbiterBridgeTask {
  /**  */
  amount: Amount;

  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  orbiterReceiverAddr?: string;

  /**  */
  orbiterSenderAddr?: string;

  /**  */
  swapCompleted?: boolean;

  /**  */
  tx?: TaskTx;
}

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

  /**  */
  progress: string;

  /**  */
  deletedAt?: Date;

  /**  */
  autoRetry: boolean;

  /**  */
  flow: flow_Flow;
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

  /**  */
  profileLabel: string;
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

export interface SetTaskSettingsRequest {
  /**  */
  wei: string;

  /**  */
  taskType: TaskType;
}

export interface SetTaskSettingsResponse {}

export interface SkipProcessTaskRequest {
  /**  */
  taskId: string;

  /**  */
  processId: string;

  /**  */
  profileId: string;
}

export interface SkipProcessTaskResponse {}

export interface SnapshotVoteTask {
  /**  */
  space: string;

  /**  */
  proposal: object;
}

export interface SpaceFiSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface StargateBridgeTask {
  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  fee?: string;

  /**  */
  txId?: string;

  /**  */
  layerZeroStatus?: string;

  /**  */
  lzscanUrl?: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  amount: Amount;

  /**  */
  tx?: TaskTx;
}

export interface StopProcessRequest {
  /**  */
  processId: string;
}

export interface StopProcessResponse {}

export interface Swap1inchTask {
  /**  */
  network: Network;

  /**  */
  fromTokenName: string;

  /**  */
  fromTokenCode: string;

  /**  */
  fromTokenAddr: string;

  /**  */
  toTokenName: string;

  /**  */
  toTokenCode: string;

  /**  */
  toTokenAddr: string;

  /**  */
  sendAll?: boolean;

  /**  */
  sendPercent?: string;

  /**  */
  sendAmount?: number;

  /**  */
  txId?: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  swapCompleted?: boolean;
}

export interface SyncSwapLPTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  a: Token;

  /**  */
  b: Token;

  /**  */
  tx?: TaskTx;

  /**  */
  add: boolean;
}

export interface SyncSwapTask {
  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  amount: Amount;

  /**  */
  amountKind: string;

  /**  */
  txId?: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  swapCompleted?: boolean;

  /**  */
  tx?: TaskTx;
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

  /**  */
  okexDepositTask?: OkexDepositTask;

  /**  */
  testNetBridgeSwapTask?: TestNetBridgeSwapTask;

  /**  */
  snapshotVoteTask?: SnapshotVoteTask;

  /**  */
  okexBinanaceTask?: OkexBinanaceTask;

  /**  */
  swap1inchTask?: Swap1inchTask;

  /**  */
  syncSwapTask?: SyncSwapTask;

  /**  */
  zkSyncOfficialBridgeToEthereumTask?: ZkSyncOfficialBridgeToEthereumTask;

  /**  */
  orbiterBridgeTask?: OrbiterBridgeTask;

  /**  */
  zkSyncOfficialBridgeFromEthereumTask?: ZkSyncOfficialBridgeFromEthereumTask;

  /**  */
  wETHTask?: WETHTask;

  /**  */
  muteioSwapTask?: MuteioSwapTask;

  /**  */
  syncSwapLPTask?: SyncSwapLPTask;

  /**  */
  maverickSwapTask?: MaverickSwapTask;

  /**  */
  spaceFiSwapTask?: SpaceFiSwapTask;

  /**  */
  velocoreSwapTask?: VelocoreSwapTask;

  /**  */
  izumiSwapTask?: IzumiSwapTask;

  /**  */
  veSyncSwapTask?: VeSyncSwapTask;

  /**  */
  ezkaliburSwapTask?: EzkaliburSwapTask;

  /**  */
  zkSwapTask?: ZkSwapTask;
}

export interface TaskTx {
  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  retryCount?: string;

  /**  */
  url?: string;

  /**  */
  network?: Network;

  /**  */
  code?: string;

  /**  */
  gasEstimated?: AmUni;

  /**  */
  gasResult?: AmUni;

  /**  */
  gasLimit?: AmUni;

  /**  */
  multiplier?: number;
}

export interface TestNetBridgeSwapTask {
  /**  */
  network: Network;

  /**  */
  minAmount: string;

  /**  */
  maxAmount: string;

  /**  */
  amount?: string;

  /**  */
  tx?: TaskTx;
}

export interface Transaction {
  /**  */
  id: string;

  /**  */
  code: string;

  /**  */
  taskId: string;

  /**  */
  profileId: string;

  /**  */
  processId: string;

  /**  */
  userId: string;

  /**  */
  createdAt: Date;

  /**  */
  network: Network;

  /**  */
  url: string;
}

export interface VeSyncSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface VelocoreSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface WETHTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  wrap: boolean;

  /**  */
  tx?: TaskTx;
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

  /**  */
  withdrawAddr?: string;

  /**  */
  useExternalAddr?: boolean;

  /**  */
  sendAllCoins?: boolean;
}

export interface ZkSwapTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;
}

export interface ZkSyncOfficialBridgeFromEthereumTask {
  /**  */
  amount: Amount;

  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  tx?: TaskTx;
}

export interface ZkSyncOfficialBridgeToEthereumTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  amountKind: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  tx?: TaskTx;
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

  /**  */
  deletedAt?: Date;
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

export enum Token {
  'USDT' = 'USDT',
  'ETH' = 'ETH',
  'USDC' = 'USDC',
  'STG' = 'STG',
  'BNB' = 'BNB',
  'MATIC' = 'MATIC',
  'AVAX' = 'AVAX',
  'veSTG' = 'veSTG',
  'WETH' = 'WETH',
  'LUSD' = 'LUSD',
  'LSD' = 'LSD',
  'MUTE' = 'MUTE',
  'MAV' = 'MAV',
  'SPACE' = 'SPACE',
  'VC' = 'VC',
  'IZI' = 'IZI'
}
