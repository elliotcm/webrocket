WEBROCKET - Distributed Web Broker for the Masses!
==================================================

WebRocket is a hybrid of MQ and WebSockets server with great support for
horizontal scalability. WebRocket is very fast and easy to use from both
sides: backend (via MQ connection) and frontend (via WebSockets).
This combination will lead you to new quality of web development and
will finally make bidirectional Web easy for everyone.

Installation
------------
The file 'INSTALL' in this directory explains how to build and install
WebRocket  on various systems, once you have unpacked or checked out the
entire WebRocket file tree.

Release history
---------------
See the file 'NEWS' for information on new features and other user-visible
changes in recent versions of WebRocket.

Reporting bugs
--------------
If you think you may have found a bug in WebRocket, please report it
using Github's internal issue tracker for this project. If you know 
how to fix the bug, please get familiar with the 'Contributing' section
of this document and don't hesitate to send a pull request.

Contributing
------------
We appreciate any kind of help in WebRocket development and testing.
If you don't know how can you help us, please check the project's
issues or 'TODO' file in the project root directory. Before you start
working on the code, and before you send us a pull request, make sure
you understand and follow the contributors guide, which you can find
in the 'CONTRIBUTE' file.

Project structure
-----------------
The main source tree contains several directories. Here's the short
explanation what you can find in there.

'docs'::
	Documentation and man pages

'distros'::
	Tools and files specific for various operating systems

'pkg/webrocket'::
	WebRocket core library

'pkg/kosmonaut'::
	WebRocket backend client library
    
'cmd/webrocket-server'::
	Tool for starting and preconfiguring single WebRocket
	server node.

'cmd/webrocket-admin'::
	Tool for ad-hoc configuration management of the running
	WebRocket server nodes.

'cmd/webrocket-monitor'::
	Tool for WebRocket server nodes and clusters monitoring.

Sponsors
--------
All the work on the project is sponsored and supported by Cubox - an
awesome dev shop from Uruguay <http://cuboxlabs.com>.

Copyright
---------
Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch> and folks at Cubox

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
