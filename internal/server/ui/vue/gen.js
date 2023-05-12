// tslint:disable
/* eslint-disable */
const {codegen} = require('swagger-axios-codegen');
const path = require('path');
// const {tr} = require("vuetify/locale");

codegen({
  methodNameMode: 'operationId',
  source: require(path.resolve(__dirname, 'swagger/profile.swagger.json')),
  outputDir: path.resolve(__dirname, 'src/generated/'),
  useCustomerRequestInstance: true,
  serviceNameSuffix: '',
  useStaticMethod: false,
  include: ['*'],
  fileName: 'profile.ts'
});

codegen({
  methodNameMode: 'operationId',
  source: require(path.resolve(__dirname, 'swagger/helper.swagger.json')),
  outputDir: path.resolve(__dirname, 'src/generated/'),
  useCustomerRequestInstance: true,
  serviceNameSuffix: '',
  useStaticMethod: false,
  include: ['*'],
  fileName: 'helper.ts'
});

codegen({
  methodNameMode: 'operationId',
  source: require(path.resolve(__dirname, 'swagger/withdraw.swagger.json')),
  outputDir: path.resolve(__dirname, 'src/generated/'),
  useCustomerRequestInstance: true,
  serviceNameSuffix: '',
  useStaticMethod: false,
  include: ['*'],
  fileName: 'withdraw.ts'
});

codegen({
  methodNameMode: 'operationId',
  source: require(path.resolve(__dirname, 'swagger/flow.swagger.json')),
  outputDir: path.resolve(__dirname, 'src/generated/'),
  useCustomerRequestInstance: true,
  serviceNameSuffix: '',
  useStaticMethod: false,
  include: ['*'],
  fileName: 'flow.ts'
});

codegen({
  methodNameMode: 'operationId',
  source: require(path.resolve(__dirname, 'swagger/process.swagger.json')),
  outputDir: path.resolve(__dirname, 'src/generated/'),
  useCustomerRequestInstance: true,
  serviceNameSuffix: '',
  useStaticMethod: false,
  include: ['*'],
  fileName: 'process.ts'
});


codegen({
  methodNameMode: 'operationId',
  source: require(path.resolve(__dirname, 'swagger/settings.swagger.json')),
  outputDir: path.resolve(__dirname, 'src/generated/'),
  useCustomerRequestInstance: true,
  serviceNameSuffix: '',
  useStaticMethod: false,
  include: ['*'],
  fileName: 'settings.ts'
});
