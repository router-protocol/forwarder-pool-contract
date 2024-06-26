#!/usr/bin/env node

const fs = require('fs');
// const rimraf = require('rimraf');

const BUILD_PATH = './build/bindings/';
const ABI_PATH = BUILD_PATH + 'abi/';
const BIN_PATH = BUILD_PATH + 'bin/';
const RUNTIME_PATH = BUILD_PATH + 'runtime/';

// Loop through all the files in the temp directory
fs.readdir('./out', function (err, files) {
    if (err) {
        console.error('Could not list the directory.', err);
        process.exit(1);
    }

    // Remove old build
    // rimraf.sync(BUILD_PATH);

    // Create empty dirs
    fs.mkdirSync(BUILD_PATH, { recursive: true });
    if (!fs.existsSync(ABI_PATH)) {
        fs.mkdirSync(ABI_PATH);
    }
    if (!fs.existsSync(BIN_PATH)) {
        fs.mkdirSync(BIN_PATH);
    }
    if (!fs.existsSync(RUNTIME_PATH)) {
        fs.mkdirSync(RUNTIME_PATH);
    }

    files.forEach(function (file, index) {
        const basename = file.split('.')[0];
        if (file.split('.')[1] == 'sol') {
            try {
                const path = './out/' + file + '/' + basename + '.json';
                let rawdata = fs.readFileSync(path);
                let contract = JSON.parse(rawdata);
                let abi = contract.abi;
                let bytecode = contract.bytecode.object;

                if (abi.length === 0) return;
                fs.writeFileSync(ABI_PATH + basename + '.abi', JSON.stringify(abi));
                fs.writeFileSync(BIN_PATH + basename + '.bin', bytecode);
                fs.writeFileSync(
                    RUNTIME_PATH + basename + '.json',
                    JSON.stringify({ abi: abi, bytecode: bytecode }),
                );
            } catch { }
        }
    });
});