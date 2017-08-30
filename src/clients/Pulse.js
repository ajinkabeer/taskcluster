// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.

import Client from '../Client';

export default class Pulse extends Client {
  constructor(options = {}) {
    super({
      ...options,
      baseUrl: 'https://pulse.taskcluster.net/v1',
      exchangePrefix: ''
    });
    
    this.overview.entryReference = {type:'function',method:'get',route:'/overview',query:[],args:[],name:'overview',stability:'experimental',title:'Rabbit Overview',description:'Get an overview of the Rabbit cluster.',output:'http://schemas.taskcluster.net/pulse/v1/rabbit-overview.json'};
    this.listNamespaces.entryReference = {type:'function',method:'get',route:'/namespaces',query:['limit','continuation'],args:[],name:'listNamespaces',stability:'experimental',title:'List Namespaces',description:'List the namespaces managed by this service.\n\nThis will list up to 1000 namespaces. If more namespaces are present a\n`continuationToken` will be returned, which can be given in the next\nrequest. For the initial request, do not provide continuation.',output:'http://schemas.taskcluster.net/pulse/v1/list-namespaces-response.json'};
    this.namespace.entryReference = {type:'function',method:'get',route:'/namespace/<namespace>',query:[],args:['namespace'],name:'namespace',stability:'experimental',title:'Get a namespace',description:'Get public information about a single namespace. This is the same information\nas returned by `listNamespaces`.',output:'http://schemas.taskcluster.net/pulse/v1/namespace.json'};
    this.claimNamespace.entryReference = {type:'function',method:'post',route:'/namespace/<namespace>',query:[],args:['namespace'],name:'claimNamespace',stability:'experimental',title:'Claim a namespace',description:'Claim a namespace, returning a username and password with access to that\nnamespace good for a short time.  Clients should call this endpoint again\nat the re-claim time given in the response, as the password will be rotated\nsoon after that time.  The namespace will expire, and any associated queues\nand exchanges will be deleted, at the given expiration time.\n\nThe `expires` and `contact` properties can be updated at any time in a reclaim\noperation.',scopes:[['pulse:namespace:<namespace>']],input:'http://schemas.taskcluster.net/pulse/v1/namespace-request.json',output:'http://schemas.taskcluster.net/pulse/v1/namespace-response.json'};
    this.ping.entryReference = {type:'function',method:'get',route:'/ping',query:[],args:[],name:'ping',stability:'stable',title:'Ping Server',description:'Respond without doing anything.\nThis endpoint is used to check that the service is up.'};
  }

  // Get an overview of the Rabbit cluster.
  overview(...args) {
    this.validateMethod(this.overview.entryReference, args);
    return this.request(this.overview.entryReference, args);
  }

  // List the namespaces managed by this service.
  // This will list up to 1000 namespaces. If more namespaces are present a
  // `continuationToken` will be returned, which can be given in the next
  // request. For the initial request, do not provide continuation.
  listNamespaces(...args) {
    this.validateMethod(this.listNamespaces.entryReference, args);
    return this.request(this.listNamespaces.entryReference, args);
  }

  // Get public information about a single namespace. This is the same information
  // as returned by `listNamespaces`.
  namespace(...args) {
    this.validateMethod(this.namespace.entryReference, args);
    return this.request(this.namespace.entryReference, args);
  }

  // Claim a namespace, returning a username and password with access to that
  // namespace good for a short time.  Clients should call this endpoint again
  // at the re-claim time given in the response, as the password will be rotated
  // soon after that time.  The namespace will expire, and any associated queues
  // and exchanges will be deleted, at the given expiration time.
  // The `expires` and `contact` properties can be updated at any time in a reclaim
  // operation.
  claimNamespace(...args) {
    this.validateMethod(this.claimNamespace.entryReference, args);
    return this.request(this.claimNamespace.entryReference, args);
  }

  // Respond without doing anything.
  // This endpoint is used to check that the service is up.
  ping(...args) {
    this.validateMethod(this.ping.entryReference, args);
    return this.request(this.ping.entryReference, args);
  }
}
