<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: google/protobuf/any.proto
//   Date: 2016-04-19 16:27:33

namespace google\protobuf {

  class Any extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $type_url = null;
    
    /**  @var string */
    public $value = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'google.protobuf.Any');

      // OPTIONAL STRING type_url = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "type_url";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL BYTES value = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "value";
      $f->type      = \DrSlump\Protobuf::TYPE_BYTES;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <type_url> has a value
     *
     * @return boolean
     */
    public function hasTypeUrl(){
      return $this->_has(1);
    }
    
    /**
     * Clear <type_url> value
     *
     * @return \google\protobuf\Any
     */
    public function clearTypeUrl(){
      return $this->_clear(1);
    }
    
    /**
     * Get <type_url> value
     *
     * @return string
     */
    public function getTypeUrl(){
      return $this->_get(1);
    }
    
    /**
     * Set <type_url> value
     *
     * @param string $value
     * @return \google\protobuf\Any
     */
    public function setTypeUrl( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <value> has a value
     *
     * @return boolean
     */
    public function hasValue(){
      return $this->_has(2);
    }
    
    /**
     * Clear <value> value
     *
     * @return \google\protobuf\Any
     */
    public function clearValue(){
      return $this->_clear(2);
    }
    
    /**
     * Get <value> value
     *
     * @return string
     */
    public function getValue(){
      return $this->_get(2);
    }
    
    /**
     * Set <value> value
     *
     * @param string $value
     * @return \google\protobuf\Any
     */
    public function setValue( $value){
      return $this->_set(2, $value);
    }
  }
}

