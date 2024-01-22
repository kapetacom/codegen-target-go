/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */
import _ from 'lodash';

export type MapUnknown = { [key: string]: any };

/**
 * Add new properties and update unchanged properties from newContent to target.
 *
 * We compare with lastContent to avoid overwriting user changes.
 */
export const addNewOrUnchanged = (target: MapUnknown, newContent: MapUnknown, lastContent: MapUnknown | null) => {
    Object.entries(newContent).forEach(([key, value]) => {
        if (lastContent) {
            if (!(key in target) && key in lastContent) {
                // User has removed the property, so we should not add it back
                return;
            }

            if (key in target && key in lastContent && !_.isEqual(target[key], lastContent[key])) {
                // User has changed the property, so we should not overwrite the change
                return;
            }
        } else if (key in target) {
            // Don't overwrite existing properties
            return;
        }

        target[key] = _.cloneDeep(value);
    });

    if (lastContent) {
        Object.entries(lastContent).forEach(([key, value]) => {
            if (!(key in newContent) && target[key] === value) {
                // Property was not changed by user, but removed in new generation, so we should remove it
                delete target[key];
            }
        });
    }
};
