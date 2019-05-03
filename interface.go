/*
 * Copyright (c) 2019.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * By gonejack@gmail.com
 */

package logger

// logger instance interface
type Logger interface {
	Debugf(tpl string, values ...interface{})
	Infof(tpl string, values ...interface{})
	Warnf(tpl string, values ...interface{})
	Errorf(tpl string, values ...interface{})
	Fatalf(tpl string, values ...interface{})
}
