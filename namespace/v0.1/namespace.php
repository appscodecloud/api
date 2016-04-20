<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: namespace.proto
//   Date: 2016-04-19 16:27:34

namespace namespace {

  class CheckRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $name = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.CheckRequest');

      // OPTIONAL STRING name = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <name> has a value
     *
     * @return boolean
     */
    public function hasName(){
      return $this->_has(1);
    }
    
    /**
     * Clear <name> value
     *
     * @return \namespace\CheckRequest
     */
    public function clearName(){
      return $this->_clear(1);
    }
    
    /**
     * Get <name> value
     *
     * @return string
     */
    public function getName(){
      return $this->_get(1);
    }
    
    /**
     * Set <name> value
     *
     * @param string $value
     * @return \namespace\CheckRequest
     */
    public function setName( $value){
      return $this->_set(1, $value);
    }
  }
}

namespace namespace {

  class CheckResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.CheckResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
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
     * @return \namespace\CheckResponse
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
     * @return \namespace\CheckResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
  }
}

namespace namespace {

  class CreateRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $name = null;
    
    /**  @var string */
    public $display_name = null;
    
    /**  @var string */
    public $email = null;
    
    /**  @var string */
    public $user_name = null;
    
    /**  @var string */
    public $password = null;
    
    /**  @var string[]  */
    public $invite_email = array();
    
    /**  @var string */
    public $user_ip = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.CreateRequest');

      // OPTIONAL STRING name = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING display_name = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "display_name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING email = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "email";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING user_name = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "user_name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING password = 5
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 5;
      $f->name      = "password";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // REPEATED STRING invite_email = 6
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 6;
      $f->name      = "invite_email";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      // OPTIONAL STRING user_ip = 7
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 7;
      $f->name      = "user_ip";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <name> has a value
     *
     * @return boolean
     */
    public function hasName(){
      return $this->_has(1);
    }
    
    /**
     * Clear <name> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearName(){
      return $this->_clear(1);
    }
    
    /**
     * Get <name> value
     *
     * @return string
     */
    public function getName(){
      return $this->_get(1);
    }
    
    /**
     * Set <name> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setName( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <display_name> has a value
     *
     * @return boolean
     */
    public function hasDisplayName(){
      return $this->_has(2);
    }
    
    /**
     * Clear <display_name> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearDisplayName(){
      return $this->_clear(2);
    }
    
    /**
     * Get <display_name> value
     *
     * @return string
     */
    public function getDisplayName(){
      return $this->_get(2);
    }
    
    /**
     * Set <display_name> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setDisplayName( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <email> has a value
     *
     * @return boolean
     */
    public function hasEmail(){
      return $this->_has(3);
    }
    
    /**
     * Clear <email> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearEmail(){
      return $this->_clear(3);
    }
    
    /**
     * Get <email> value
     *
     * @return string
     */
    public function getEmail(){
      return $this->_get(3);
    }
    
    /**
     * Set <email> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setEmail( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <user_name> has a value
     *
     * @return boolean
     */
    public function hasUserName(){
      return $this->_has(4);
    }
    
    /**
     * Clear <user_name> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearUserName(){
      return $this->_clear(4);
    }
    
    /**
     * Get <user_name> value
     *
     * @return string
     */
    public function getUserName(){
      return $this->_get(4);
    }
    
    /**
     * Set <user_name> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setUserName( $value){
      return $this->_set(4, $value);
    }
    
    /**
     * Check if <password> has a value
     *
     * @return boolean
     */
    public function hasPassword(){
      return $this->_has(5);
    }
    
    /**
     * Clear <password> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearPassword(){
      return $this->_clear(5);
    }
    
    /**
     * Get <password> value
     *
     * @return string
     */
    public function getPassword(){
      return $this->_get(5);
    }
    
    /**
     * Set <password> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setPassword( $value){
      return $this->_set(5, $value);
    }
    
    /**
     * Check if <invite_email> has a value
     *
     * @return boolean
     */
    public function hasInviteEmail(){
      return $this->_has(6);
    }
    
    /**
     * Clear <invite_email> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearInviteEmail(){
      return $this->_clear(6);
    }
    
    /**
     * Get <invite_email> value
     *
     * @param int $idx
     * @return string
     */
    public function getInviteEmail($idx = NULL){
      return $this->_get(6, $idx);
    }
    
    /**
     * Set <invite_email> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setInviteEmail( $value, $idx = NULL){
      return $this->_set(6, $value, $idx);
    }
    
    /**
     * Get all elements of <invite_email>
     *
     * @return string[]
     */
    public function getInviteEmailList(){
     return $this->_get(6);
    }
    
    /**
     * Add a new element to <invite_email>
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function addInviteEmail( $value){
     return $this->_add(6, $value);
    }
    
    /**
     * Check if <user_ip> has a value
     *
     * @return boolean
     */
    public function hasUserIp(){
      return $this->_has(7);
    }
    
    /**
     * Clear <user_ip> value
     *
     * @return \namespace\CreateRequest
     */
    public function clearUserIp(){
      return $this->_clear(7);
    }
    
    /**
     * Get <user_ip> value
     *
     * @return string
     */
    public function getUserIp(){
      return $this->_get(7);
    }
    
    /**
     * Set <user_ip> value
     *
     * @param string $value
     * @return \namespace\CreateRequest
     */
    public function setUserIp( $value){
      return $this->_set(7, $value);
    }
  }
}

namespace namespace {

  class CreateResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.CreateResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
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
     * @return \namespace\CreateResponse
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
     * @return \namespace\CreateResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
  }
}

namespace namespace {

  class LogRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $name = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.LogRequest');

      // OPTIONAL STRING name = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <name> has a value
     *
     * @return boolean
     */
    public function hasName(){
      return $this->_has(1);
    }
    
    /**
     * Clear <name> value
     *
     * @return \namespace\LogRequest
     */
    public function clearName(){
      return $this->_clear(1);
    }
    
    /**
     * Get <name> value
     *
     * @return string
     */
    public function getName(){
      return $this->_get(1);
    }
    
    /**
     * Set <name> value
     *
     * @param string $value
     * @return \namespace\LogRequest
     */
    public function setName( $value){
      return $this->_set(1, $value);
    }
  }
}

namespace namespace {

  class LogResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    
    /**  @var \namespace\Log[]  */
    public $logs = array();
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.LogResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
      $descriptor->addField($f);

      // REPEATED MESSAGE logs = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "logs";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $f->reference = '\namespace\Log';
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
     * @return \namespace\LogResponse
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
     * @return \namespace\LogResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <logs> has a value
     *
     * @return boolean
     */
    public function hasLogs(){
      return $this->_has(2);
    }
    
    /**
     * Clear <logs> value
     *
     * @return \namespace\LogResponse
     */
    public function clearLogs(){
      return $this->_clear(2);
    }
    
    /**
     * Get <logs> value
     *
     * @param int $idx
     * @return \namespace\Log
     */
    public function getLogs($idx = NULL){
      return $this->_get(2, $idx);
    }
    
    /**
     * Set <logs> value
     *
     * @param \namespace\Log $value
     * @return \namespace\LogResponse
     */
    public function setLogs(\namespace\Log $value, $idx = NULL){
      return $this->_set(2, $value, $idx);
    }
    
    /**
     * Get all elements of <logs>
     *
     * @return \namespace\Log[]
     */
    public function getLogsList(){
     return $this->_get(2);
    }
    
    /**
     * Add a new element to <logs>
     *
     * @param \namespace\Log $value
     * @return \namespace\LogResponse
     */
    public function addLogs(\namespace\Log $value){
     return $this->_add(2, $value);
    }
  }
}

namespace namespace {

  class Log extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $info = null;
    
    /**  @var string */
    public $details = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.Log');

      // OPTIONAL STRING info = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "info";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING details = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "details";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <info> has a value
     *
     * @return boolean
     */
    public function hasInfo(){
      return $this->_has(1);
    }
    
    /**
     * Clear <info> value
     *
     * @return \namespace\Log
     */
    public function clearInfo(){
      return $this->_clear(1);
    }
    
    /**
     * Get <info> value
     *
     * @return string
     */
    public function getInfo(){
      return $this->_get(1);
    }
    
    /**
     * Set <info> value
     *
     * @param string $value
     * @return \namespace\Log
     */
    public function setInfo( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <details> has a value
     *
     * @return boolean
     */
    public function hasDetails(){
      return $this->_has(2);
    }
    
    /**
     * Clear <details> value
     *
     * @return \namespace\Log
     */
    public function clearDetails(){
      return $this->_clear(2);
    }
    
    /**
     * Get <details> value
     *
     * @return string
     */
    public function getDetails(){
      return $this->_get(2);
    }
    
    /**
     * Set <details> value
     *
     * @param string $value
     * @return \namespace\Log
     */
    public function setDetails( $value){
      return $this->_set(2, $value);
    }
  }
}

namespace namespace {

  class StatusRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $name = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.StatusRequest');

      // OPTIONAL STRING name = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <name> has a value
     *
     * @return boolean
     */
    public function hasName(){
      return $this->_has(1);
    }
    
    /**
     * Clear <name> value
     *
     * @return \namespace\StatusRequest
     */
    public function clearName(){
      return $this->_clear(1);
    }
    
    /**
     * Get <name> value
     *
     * @return string
     */
    public function getName(){
      return $this->_get(1);
    }
    
    /**
     * Set <name> value
     *
     * @param string $value
     * @return \namespace\StatusRequest
     */
    public function setName( $value){
      return $this->_set(1, $value);
    }
  }
}

namespace namespace {

  class StatusResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'namespace.StatusResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
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
     * @return \namespace\StatusResponse
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
     * @return \namespace\StatusResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
  }
}

namespace namespace {

  class NamespaceClient extends \Grpc\BaseStub {

    public function __construct($hostname, $opts) {
      parent::__construct($hostname, $opts);
    }
    /**
     * @param namespace\CheckRequest $input
     */
    public function Check(\namespace\CheckRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/namespace.Namespace/Check', $argument, '\namespace\CheckResponse::deserialize', $metadata, $options);
    }
    /**
     * @param namespace\CreateRequest $input
     */
    public function Create(\namespace\CreateRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/namespace.Namespace/Create', $argument, '\namespace\CreateResponse::deserialize', $metadata, $options);
    }
    /**
     * @param namespace\StatusRequest $input
     */
    public function Status(\namespace\StatusRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/namespace.Namespace/Status', $argument, '\namespace\StatusResponse::deserialize', $metadata, $options);
    }
    /**
     * @param namespace\LogRequest $input
     */
    public function Log(\namespace\LogRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/namespace.Namespace/Log', $argument, '\namespace\LogResponse::deserialize', $metadata, $options);
    }
  }
}
