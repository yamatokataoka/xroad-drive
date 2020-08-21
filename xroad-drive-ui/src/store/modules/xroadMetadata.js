export default {
  namespaced: true,
  state() {
    return {
      providers: [
        {
          id: 'CS:ORG:1111:Company1Provider',
          name: 'Company 1'
        }, {
          id: 'CS:ORG:1112:Company2Provider',
          name: 'Company 2'
        }, {
          id: 'CS:ORG:1113:Company3Provider',
          name: 'Company 3'
        }
      ],
      clients: [
        {
          id: 'CS:ORG:1111:Company1Provider',
          name: 'Company 1'
        }, {
          id: 'CS:ORG:1114:Company4Provider',
          name: 'Company 4'
        }, {
          id: 'CS:ORG:1115:Company5Provider',
          name: 'Company 5'
        }
      ]
    };
  },
  getters: {
    providersForNav(state) {
      return parentName => {
        let providersForNav = [];
        for (let i = 0; i < state.providers.length; i++) {
          providersForNav[i] = {...state.providers[i]};
          providersForNav[i].id = 'shared-with-us' + ':' + state.providers[i].id;
          providersForNav[i].to = { name: parentName, params: { id: state.providers[i].id }};
          providersForNav[i].parent = parentName;
        }
        return providersForNav;
      }
    },
    clientsForNav(state) {
      return parentName => {
        let clientsForNav = [];
        for (let i = 0; i < state.clients.length; i++) {
          clientsForNav[i] = {...state.clients[i]};
          clientsForNav[i].id = 'shared-with-others' + ':' + state.clients[i].id;
          clientsForNav[i].to = { name: parentName, params: { id: state.clients[i].id }};
          clientsForNav[i].parent = parentName;
        }
        return clientsForNav;
      }
    }
  }
};
