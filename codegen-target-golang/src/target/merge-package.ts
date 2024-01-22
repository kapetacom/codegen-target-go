/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */
import { GeneratedFile, SourceFile } from '@kapeta/codegen';
import { addNewOrUnchanged } from './helpers';

export const mergePackageJson = (
    sourceFile: SourceFile,
    newFile: GeneratedFile,
    lastFile: GeneratedFile | null
): GeneratedFile => {
    // We can merge the dependencies and scripts into existing package.json without overwriting
    // the existing user adjusted content

    const target = JSON.parse(sourceFile.content);
    const newContent = JSON.parse(newFile.content);
    const lastContent = lastFile ? JSON.parse(lastFile.content) : null;
    if (!target.dependencies) {
        target.dependencies = {};
    }

    if (!target.devDependencies) {
        target.devDependencies = {};
    }

    addNewOrUnchanged(target.dependencies, newContent.dependencies, lastContent?.dependencies ?? null);
    addNewOrUnchanged(target.devDependencies, newContent.devDependencies, lastContent?.devDependencies ?? null);

    if (!target.scripts) {
        target.scripts = {};
    }
    addNewOrUnchanged(target.scripts, newContent.scripts, lastContent?.scripts ?? null);
    addNewOrUnchanged(target, newContent, lastContent);

    return {
        ...newFile,
        content: JSON.stringify(target, null, 4),
    };
};
