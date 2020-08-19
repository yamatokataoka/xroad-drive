export default {
  namespaced: true,
  state() {
    return {
      providers: [
        {
          id: 'shared-with-us:CS:ORG:1111:Company1Provider',
          name: 'Company 1',
          type: 'provider',
          to: '/shared-with-us/CS:ORG:1111:Company1Provider',
        }, {
          id: 'shared-with-us:CS:ORG:1112:Company2Provider',
          name: 'Company 2',
          type: 'provider',
          to: '/shared-with-us/CS:ORG:1112:Company2Provider',
        }, {
          id: 'shared-with-us:CS:ORG:1113:Company3Provider',
          name: 'Company 3',
          type: 'provider',
          to: '/shared-with-us/CS:ORG:1113:Company3Provider',
        }
      ],
      clients: [
        {
          id: 'shared-with-others:CS:ORG:1111:Company1Provider',
          name: 'Company 1',
          type: 'client',
          to: '/shared-with-others/CS:ORG:1111:Company1Provider',
        }, {
          id: 'shared-with-others:CS:ORG:1114:Company4Provider',
          name: 'Company 4',
          type: 'client',
          to: '/shared-with-others/CS:ORG:1114:Company4Provider',
        }, {
          id: 'shared-with-others:CS:ORG:1115:Company5Provider',
          name: 'Company 5',
          type: 'client',
          to: '/shared-with-others/CS:ORG:1115:Company5Provider',
        }
      ]
    };
  }
};
