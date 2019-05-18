#Create Database:

CREATE Database edushop;

#shop_products
CREATE TABLE `shop_products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'product name',
  `model` varchar(255) NOT NULL COMMENT 'product type or model',
  `price`  DECIMAL(10,3)  NOT NULL COMMENT 'price of product',
  `desc` varchar(255) DEFAULT '' COMMENT 'product description',
  `image_url` varchar(255) DEFAULT '' COMMENT 'image of the product for introduction',
  `video_url` varchar(255) DEFAULT '' COMMENT 'video of the product for introduction',
  `capacity` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'create by shopper',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `modified_by` varchar(255) DEFAULT '' COMMENT 'modifiy by seller',
  `labels` varchar(255) DEFAULT '' COMMENT 'labels for searching',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'on/off',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='product managements';

#shop_users

CREATE TABLE `shop_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT 'username',
  `alias` varchar(100) DEFAULT '' COMMENT 'alias',
  `password` varchar(100) DEFAULT '' COMMENT 'password',
  `mobile` varchar(100) DEFAULT '' COMMENT 'mobile',
  `email` varchar(100) DEFAULT '' COMMENT 'email',
  `address_ids` varchar(100) DEFAULT '' COMMENT 'address ids',
  `payment_method_1` varchar(100) DEFAULT '' COMMENT 'ali/weichat/bank/credit',
  `payment_method_2` varchar(100) DEFAULT '' COMMENT 'ali/weichat/bank/credit',
  `payment_method_3` varchar(100) DEFAULT '' COMMENT 'ali/weichat/bank/credit',
  `desc` varchar(255) DEFAULT '' COMMENT 'product description',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'create by',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `modified_by` varchar(255) DEFAULT '' COMMENT 'modifiy by seller',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'enable/disable',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user managements';

# Addresses

CREATE TABLE `shop_user_addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned DEFAULT '0' COMMENT 'user_id',
  `street` varchar(100) DEFAULT '' COMMENT 'password',
  `city` varchar(100) DEFAULT '' COMMENT 'mobile',
  `state` varchar(100) DEFAULT '' COMMENT 'email',
  `country` varchar(20) DEFAULT '' COMMENT 'address',
  `short_addr` varchar(255) DEFAULT '' COMMENT 'product description',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `labels` varchar(255) DEFAULT '' COMMENT 'labels for searching',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user addresses managements';

# shop_cart
CREATE TABLE `shop_cart`(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned COMMENT 'buy user',
  `product_id` int(10) unsigned COMMENT 'buy user',
  `amount` int(10) unsigned DEFAULT '0' COMMENT 'total amount of the products',
  `total_price`  DECIMAL(10,3)  DEFAULT '0.00' COMMENT 'price of product',
  `desc` varchar(255) DEFAULT 'nothing to show' COMMENT 'product description',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'create by',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `modified_by` varchar(255) DEFAULT '' COMMENT 'modifiy by seller',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='shop char managements';

# shtop orders 
CREATE TABLE `shop_orders` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned COMMENT 'product id for order',
  `user_id` int(10) unsigned  COMMENT 'buy user',
  `payment_method_id` int(10) unsigned COMMENT 'ali/weichat/bank/credit',
  `amount` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `total_price`  DECIMAL(10,3)  DEFAULT '0.00' COMMENT 'price of product',
  `desc` varchar(255) DEFAULT '' COMMENT 'product description',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'create by',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `modified_by` varchar(255) DEFAULT '' COMMENT 'modifiy by seller',
  `labels` varchar(255) DEFAULT '' COMMENT 'labels for searching',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'on/off',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user orders managements';

#shop payment
CREATE TABLE `shop_payments`(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned COMMENT 'buy user',
  `payment_method_id` tinyint(3) unsigned DEFAULT '1' COMMENT 'ali/weichat/bank/credit',
  `amount_price`  DECIMAL(10,3)  DEFAULT '0.00' COMMENT 'price of product',
  `desc` varchar(255) DEFAULT '' COMMENT 'product description',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'create by',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `modified_by` varchar(255) DEFAULT '' COMMENT 'modifiy by seller',
  `labels` varchar(255) DEFAULT '' COMMENT 'labels for searching',
  `status` varchar(255) DEFAULT '' COMMENT 'closed/processing/pending',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='shop char managements';

# m2m payment-orders
CREATE TABLE `shop_payments_orders_id`(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `payment_id` int(10) unsigned COMMENT 'payment id for m to m relationship',
  `order_id` int(10) unsigned COMMENT 'order_id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT=' many to many payments and orders managements'

# shop payment methods
CREATE TABLE `shop_payment_methods`(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `payment_method` varchar(100) DEFAULT 'Weichat' COMMENT 'ali/weichat/bank/credit',
  `payment_method_api` varchar(255) DEFAULT 'Weichat' COMMENT 'ali/weichat/bank/credit',
  `desc` varchar(255) DEFAULT '' COMMENT 'product description',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'create by',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modify time',
  `modified_by` varchar(255) DEFAULT '' COMMENT 'modifiy by seller',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'on/off',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='shop char managements';