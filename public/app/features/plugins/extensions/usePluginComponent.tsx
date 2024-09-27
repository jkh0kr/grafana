import { useMemo } from 'react';
import { useObservable } from 'react-use';

import { UsePluginComponentResult } from '@grafana/runtime';

import { useExposedComponentsRegistry } from './ExtensionRegistriesContext';
import { log } from './logs/log';
import { wrapWithPluginContext } from './utils';

// Returns a component exposed by a plugin.
// (Exposed components can be defined in plugins by calling .exposeComponent() on the AppPlugin instance.)
export function usePluginComponent<Props extends object = {}>(id: string): UsePluginComponentResult<Props> {
  const registry = useExposedComponentsRegistry();
  const registryState = useObservable(registry.asObservable());

  return useMemo(() => {
    if (!registryState?.[id]) {
      return {
        isLoading: false,
        component: null,
      };
    }

    const registryItem = registryState[id];
    const componentLog = log.child({
      title: registryItem.title,
      description: registryItem.description,
      pluginId: registryItem.pluginId,
    });

    return {
      isLoading: false,
      component: wrapWithPluginContext(registryItem.pluginId, registryItem.component, componentLog),
    };
  }, [id, registryState]);
}
