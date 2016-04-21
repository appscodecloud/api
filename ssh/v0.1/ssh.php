<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: ssh.proto
//   Date: 2016-04-19 16:27:37

namespace ssh {

  class SecureShellGetRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $jenkins_namespace = null;
    
    /**  @var string */
    public $cluster_namespace = null;
    
    /**  @var string */
    public $cluster_name = null;
    
    /**  @var string */
    public $instance_name = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'ssh.SecureShellGetRequest');

      // OPTIONAL STRING jenkins_namespace = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "jenkins_namespace";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING cluster_namespace = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "cluster_namespace";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING cluster_name = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "cluster_name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING instance_name = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "instance_name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <jenkins_namespace> has a value
     *
     * @return boolean
     */
    public function hasJenkinsNamespace(){
      return $this->_has(1);
    }
    
    /**
     * Clear <jenkins_namespace> value
     *
     * @return \ssh\SecureShellGetRequest
     */
    public function clearJenkinsNamespace(){
      return $this->_clear(1);
    }
    
    /**
     * Get <jenkins_namespace> value
     *
     * @return string
     */
    public function getJenkinsNamespace(){
      return $this->_get(1);
    }
    
    /**
     * Set <jenkins_namespace> value
     *
     * @param string $value
     * @return \ssh\SecureShellGetRequest
     */
    public function setJenkinsNamespace( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <cluster_namespace> has a value
     *
     * @return boolean
     */
    public function hasClusterNamespace(){
      return $this->_has(2);
    }
    
    /**
     * Clear <cluster_namespace> value
     *
     * @return \ssh\SecureShellGetRequest
     */
    public function clearClusterNamespace(){
      return $this->_clear(2);
    }
    
    /**
     * Get <cluster_namespace> value
     *
     * @return string
     */
    public function getClusterNamespace(){
      return $this->_get(2);
    }
    
    /**
     * Set <cluster_namespace> value
     *
     * @param string $value
     * @return \ssh\SecureShellGetRequest
     */
    public function setClusterNamespace( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <cluster_name> has a value
     *
     * @return boolean
     */
    public function hasClusterName(){
      return $this->_has(3);
    }
    
    /**
     * Clear <cluster_name> value
     *
     * @return \ssh\SecureShellGetRequest
     */
    public function clearClusterName(){
      return $this->_clear(3);
    }
    
    /**
     * Get <cluster_name> value
     *
     * @return string
     */
    public function getClusterName(){
      return $this->_get(3);
    }
    
    /**
     * Set <cluster_name> value
     *
     * @param string $value
     * @return \ssh\SecureShellGetRequest
     */
    public function setClusterName( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <instance_name> has a value
     *
     * @return boolean
     */
    public function hasInstanceName(){
      return $this->_has(4);
    }
    
    /**
     * Clear <instance_name> value
     *
     * @return \ssh\SecureShellGetRequest
     */
    public function clearInstanceName(){
      return $this->_clear(4);
    }
    
    /**
     * Get <instance_name> value
     *
     * @return string
     */
    public function getInstanceName(){
      return $this->_get(4);
    }
    
    /**
     * Set <instance_name> value
     *
     * @param string $value
     * @return \ssh\SecureShellGetRequest
     */
    public function setInstanceName( $value){
      return $this->_set(4, $value);
    }
  }
}

namespace ssh {

  class SecureShellGetResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    
    /**  @var \ssh\SSHKey */
    public $ssh_key = null;
    
    /**  @var string */
    public $instance_addr = null;
    
    /**  @var int */
    public $instance_port = null;
    
    /**  @var string */
    public $user = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'ssh.SecureShellGetResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
      $descriptor->addField($f);

      // OPTIONAL MESSAGE ssh_key = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "ssh_key";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\ssh\SSHKey';
      $descriptor->addField($f);

      // OPTIONAL STRING instance_addr = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "instance_addr";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL INT32 instance_port = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "instance_port";
      $f->type      = \DrSlump\Protobuf::TYPE_INT32;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING user = 5
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 5;
      $f->name      = "user";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <status> has a value
     *
     * @return boolean
     */
    public function hasStatus(){
      return $this->_has(1);
    }
    
    /**
     * Clear <status> value
     *
     * @return \ssh\SecureShellGetResponse
     */
    public function clearStatus(){
      return $this->_clear(1);
    }
    
    /**
     * Get <status> value
     *
     * @return \dtypes\Status
     */
    public function getStatus(){
      return $this->_get(1);
    }
    
    /**
     * Set <status> value
     *
     * @param \dtypes\Status $value
     * @return \ssh\SecureShellGetResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <ssh_key> has a value
     *
     * @return boolean
     */
    public function hasSshKey(){
      return $this->_has(2);
    }
    
    /**
     * Clear <ssh_key> value
     *
     * @return \ssh\SecureShellGetResponse
     */
    public function clearSshKey(){
      return $this->_clear(2);
    }
    
    /**
     * Get <ssh_key> value
     *
     * @return \ssh\SSHKey
     */
    public function getSshKey(){
      return $this->_get(2);
    }
    
    /**
     * Set <ssh_key> value
     *
     * @param \ssh\SSHKey $value
     * @return \ssh\SecureShellGetResponse
     */
    public function setSshKey(\ssh\SSHKey $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <instance_addr> has a value
     *
     * @return boolean
     */
    public function hasInstanceAddr(){
      return $this->_has(3);
    }
    
    /**
     * Clear <instance_addr> value
     *
     * @return \ssh\SecureShellGetResponse
     */
    public function clearInstanceAddr(){
      return $this->_clear(3);
    }
    
    /**
     * Get <instance_addr> value
     *
     * @return string
     */
    public function getInstanceAddr(){
      return $this->_get(3);
    }
    
    /**
     * Set <instance_addr> value
     *
     * @param string $value
     * @return \ssh\SecureShellGetResponse
     */
    public function setInstanceAddr( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <instance_port> has a value
     *
     * @return boolean
     */
    public function hasInstancePort(){
      return $this->_has(4);
    }
    
    /**
     * Clear <instance_port> value
     *
     * @return \ssh\SecureShellGetResponse
     */
    public function clearInstancePort(){
      return $this->_clear(4);
    }
    
    /**
     * Get <instance_port> value
     *
     * @return int
     */
    public function getInstancePort(){
      return $this->_get(4);
    }
    
    /**
     * Set <instance_port> value
     *
     * @param int $value
     * @return \ssh\SecureShellGetResponse
     */
    public function setInstancePort( $value){
      return $this->_set(4, $value);
    }
    
    /**
     * Check if <user> has a value
     *
     * @return boolean
     */
    public function hasUser(){
      return $this->_has(5);
    }
    
    /**
     * Clear <user> value
     *
     * @return \ssh\SecureShellGetResponse
     */
    public function clearUser(){
      return $this->_clear(5);
    }
    
    /**
     * Get <user> value
     *
     * @return string
     */
    public function getUser(){
      return $this->_get(5);
    }
    
    /**
     * Set <user> value
     *
     * @param string $value
     * @return \ssh\SecureShellGetResponse
     */
    public function setUser( $value){
      return $this->_set(5, $value);
    }
  }
}

namespace ssh {

  class SSHKey extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $public_key = null;
    
    /**  @var string */
    public $private_key = null;
    
    /**  @var string */
    public $aws_fingerprint = null;
    
    /**  @var string */
    public $openssh_fingerprint = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'ssh.SSHKey');

      // OPTIONAL BYTES public_key = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "public_key";
      $f->type      = \DrSlump\Protobuf::TYPE_BYTES;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL BYTES private_key = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "private_key";
      $f->type      = \DrSlump\Protobuf::TYPE_BYTES;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING aws_fingerprint = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "aws_fingerprint";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING openssh_fingerprint = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "openssh_fingerprint";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <public_key> has a value
     *
     * @return boolean
     */
    public function hasPublicKey(){
      return $this->_has(1);
    }
    
    /**
     * Clear <public_key> value
     *
     * @return \ssh\SSHKey
     */
    public function clearPublicKey(){
      return $this->_clear(1);
    }
    
    /**
     * Get <public_key> value
     *
     * @return string
     */
    public function getPublicKey(){
      return $this->_get(1);
    }
    
    /**
     * Set <public_key> value
     *
     * @param string $value
     * @return \ssh\SSHKey
     */
    public function setPublicKey( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <private_key> has a value
     *
     * @return boolean
     */
    public function hasPrivateKey(){
      return $this->_has(2);
    }
    
    /**
     * Clear <private_key> value
     *
     * @return \ssh\SSHKey
     */
    public function clearPrivateKey(){
      return $this->_clear(2);
    }
    
    /**
     * Get <private_key> value
     *
     * @return string
     */
    public function getPrivateKey(){
      return $this->_get(2);
    }
    
    /**
     * Set <private_key> value
     *
     * @param string $value
     * @return \ssh\SSHKey
     */
    public function setPrivateKey( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <aws_fingerprint> has a value
     *
     * @return boolean
     */
    public function hasAwsFingerprint(){
      return $this->_has(3);
    }
    
    /**
     * Clear <aws_fingerprint> value
     *
     * @return \ssh\SSHKey
     */
    public function clearAwsFingerprint(){
      return $this->_clear(3);
    }
    
    /**
     * Get <aws_fingerprint> value
     *
     * @return string
     */
    public function getAwsFingerprint(){
      return $this->_get(3);
    }
    
    /**
     * Set <aws_fingerprint> value
     *
     * @param string $value
     * @return \ssh\SSHKey
     */
    public function setAwsFingerprint( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <openssh_fingerprint> has a value
     *
     * @return boolean
     */
    public function hasOpensshFingerprint(){
      return $this->_has(4);
    }
    
    /**
     * Clear <openssh_fingerprint> value
     *
     * @return \ssh\SSHKey
     */
    public function clearOpensshFingerprint(){
      return $this->_clear(4);
    }
    
    /**
     * Get <openssh_fingerprint> value
     *
     * @return string
     */
    public function getOpensshFingerprint(){
      return $this->_get(4);
    }
    
    /**
     * Set <openssh_fingerprint> value
     *
     * @param string $value
     * @return \ssh\SSHKey
     */
    public function setOpensshFingerprint( $value){
      return $this->_set(4, $value);
    }
  }
}

namespace ssh {

  class SecureShellClient extends \Grpc\BaseStub {

    public function __construct($hostname, $opts) {
      parent::__construct($hostname, $opts);
    }
    /**
     * @param ssh\SecureShellGetRequest $input
     */
    public function Get(\ssh\SecureShellGetRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/ssh.SecureShell/Get', $argument, '\ssh\SecureShellGetResponse::deserialize', $metadata, $options);
    }
  }
}