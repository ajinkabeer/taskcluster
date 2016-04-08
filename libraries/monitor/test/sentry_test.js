suite('Sentry', () => {
  let assert = require('assert');
  let monitoring = require('../');
  let debug = require('debug')('test');
  let nock = require('nock');
  let authmock = require('./authmock');

  let monitor = null;

  suiteSetup(async () => {
    authmock.setup();

    monitor = await monitoring({
      project: 'tc-lib-monitor',
      credentials: {clientId: 'test-client', accessToken: 'test'},
      patchGlobal: false,
    });
  });

  suiteTeardown(() => {
    authmock.teardown();
  });

  test('should create sentry error', async function (done) {

    nock('https://app.getsentry.com')
      .filteringRequestBody(/.*/, '*')
      .post('/api/12345/store/', '*')
      .reply(200, () => {
        done();
      });

    monitor.sentry.client.on('error', done);

    await monitor.reportError('create sentry error test');
  });

});