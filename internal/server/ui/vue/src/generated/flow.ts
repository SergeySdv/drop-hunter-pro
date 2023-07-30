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
      let url = basePath + '/api/gw/v1/flow/create';

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
      let url = basePath + '/api/gw/v1/flow/delete';

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
      let url = basePath + '/api/gw/v1/flow/get';

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
      let url = basePath + '/api/gw/v1/flow/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceUpdateFlow(
    params: {
      /**  */
      body: UpdateFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UpdateFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/update';

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

  /**  */
  randomDuration?: string;
}

export interface DeleteFlowRequest {
  /**  */
  id: string;
}

export interface DeleteFlowResponse {}

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

export interface GetFlowRequest {
  /**  */
  id: string;
}

export interface GetFlowResponse {
  /**  */
  flow: flow_Flow;
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

export interface ListFlowRequest {}

export interface ListFlowResponse {
  /**  */
  flows: flow_Flow[];
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

export interface UpdateFlowRequest {
  /**  */
  flow: flow_Flow;
}

export interface UpdateFlowResponse {
  /**  */
  flow: flow_Flow;
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
