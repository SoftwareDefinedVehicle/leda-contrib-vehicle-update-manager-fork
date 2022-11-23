// Copyright (c) 2022 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0
//
// SPDX-License-Identifier: Apache-2.0

package main

type runLock struct {
	name string
	file int
}

func newRunLock(name string) (*runLock, error) {
	l := &runLock{
		name: name,
	}
	var err error
	l.file, err = openFile(l.name)
	defer closeFile(l.file)
	return l, err
}

func (l *runLock) Unlock() {
	closeFile(l.file)
}

func (l *runLock) TryLock() (err error) {
	l.file, err = openFile(l.name)
	if err == nil {
		err = lockFile(l.file)
		if err != nil {
			closeFile(l.file)
		}
	}
	return
}
